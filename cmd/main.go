package main

import (
	"backgammon/app/auth"
	"backgammon/config"
	"backgammon/infrastructure/dal/migrations"
	"backgammon/infrastructure/dal/ram_user_storage"
	"backgammon/infrastructure/dal/temp_session_storage"
	"backgammon/infrastructure/dal/user_storage_pgsql"
	"backgammon/infrastructure/handlers"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	bindata "github.com/golang-migrate/migrate/v4/source/go_bindata"
	"github.com/labstack/echo/v4"
)

func main() {
	serverConfig := config.NewServerConfig()

	database := user_storage_pgsql.NewUserDataStoragePGSQL(serverConfig)
	defer database.CloseDatabaseConnection()

	s:=bindata.Resource(migrations.AssetNames(), migrations.Asset)
	runDBMigrate(fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		serverConfig.Database.User,
		serverConfig.Database.Password,
		serverConfig.Database.Host,
		serverConfig.Database.Port,
		serverConfig.Database.Name), s)

	//userStorage := user_storage_pgsql.NewUserDataStoragePGSQL(serverConfig)
	userStorage := ram_user_storage.NewUserStorageRAM()
	mainSessionStorage := temp_session_storage.NewMainSessionStorage()

	tokenGenerator := auth.NewTokenGeneratorFlex(serverConfig)

	userAuthService := auth.NewUserAuthService(userStorage, mainSessionStorage, serverConfig, tokenGenerator)
	userWebSocketManageService := auth.NewWebSocketManageService(mainSessionStorage)

	userAuthHandler := handlers.NewUserAuthHandler(userAuthService)
	lobbyHandler := handlers.NewLobbyHandler(userAuthService)
	webSocketHandler := handlers.NewWebSocketHandler(userAuthService, userWebSocketManageService)

	e := echo.New()

	e.POST("/register", userAuthHandler.Register)
	e.POST("/authorize", userAuthHandler.Authorize)
	e.GET("/rooms", lobbyHandler.GetRoomsInfo)
	e.GET("/ws", webSocketHandler.Handle)

	e.Logger.Fatal(e.Start(serverConfig.Server.Port))
}

func runDBMigrate(dsn string, source *bindata.AssetSource)  {
	d, err := bindata.WithInstance(source)
	if err != nil {
		panic(err)
	}

	m, err := migrate.NewWithSourceInstance("go-bindata", d, dsn)
	if err != nil {
		panic(err)
	}

	if err = m.Up(); err != nil {
		if err == migrate.ErrNoChange {
			fmt.Println(err)
		} else {
			panic(err)
		}
	}
}
