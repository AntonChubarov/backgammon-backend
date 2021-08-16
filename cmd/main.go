package main

import (
	"backgammon/app"
	"backgammon/config"
	"backgammon/infrastructure/dal"
	"backgammon/infrastructure/dal/migrations"
	"backgammon/infrastructure/handlers"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	bindata "github.com/golang-migrate/migrate/v4/source/go_bindata"
	"github.com/labstack/echo/v4"
)

func main() {
	serverConfig := config.NewServerConfig()

	database := dal.NewDatabaseConnector(serverConfig)
	defer database.CloseDatabaseConnection()

	s:=bindata.Resource(migrations.AssetNames(), migrations.Asset)
	runDBMigrate(fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		serverConfig.Database.User,
		serverConfig.Database.Password,
		serverConfig.Database.Host,
		serverConfig.Database.Port,
		serverConfig.Database.Name), s)

	storage := dal.NewDatabaseConnector(serverConfig)
	mainSessionStorage := dal.NewMainSessionStorage()

	userAuthService := app.NewUserAuthService(storage, mainSessionStorage, serverConfig)
	userAuthHandler := handlers.NewUserAuthHandler(userAuthService)

	webSocket := handlers.NewWebSocketHandler()

	e := echo.New()
	e.GET("/", func(c echo.Context) error {return nil})
	e.GET("/ws", webSocket.Handle)
	e.POST("/register", userAuthHandler.Register)
	e.POST("/login", userAuthHandler.Authorize)

	e.Logger.Fatal(e.Start(serverConfig.Host.ServerStartPort))
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
