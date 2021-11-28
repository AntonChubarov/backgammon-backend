package main

import (
	"backgammon/app/service"
	"backgammon/config"
	"backgammon/infrastructure/repository"
	"backgammon/infrastructure/repository/migrations"
	"backgammon/infrastructure/webserver"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	bindata "github.com/golang-migrate/migrate/v4/source/go_bindata"
	"github.com/jbrodriguez/mlog"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetupAndRun() {
	mlog.Start(mlog.LevelInfo, "")

	conf := config.NewServerConfig()
	mlog.Info("server config loaded")

	database := repository.NewUserDataStoragePGSQL(conf)
	defer database.CloseDatabaseConnection()

	s := bindata.Resource(migrations.AssetNames(), migrations.Asset)
	runDBMigrate(fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		conf.Database.User,
		conf.Database.Password,
		conf.Database.Host,
		conf.Database.Port,
		conf.Database.Name), s)

	//userStorage := user_storage_pgsql.NewUserDataStoragePGSQL(conf)
	userStorage := repository.NewUserStorageRAM()
	mlog.Info("user storage initialized")

	mainSessionStorage := repository.NewSessionStorageRam()
	mlog.Info("session storage initialized")

	tokenGenerator := service.NewTokenGeneratorFlex(conf)

	userAuthService := service.NewUserAuthService(userStorage, mainSessionStorage, tokenGenerator)
	//userWebSocketManageService := service.NewWebSocketManageService(mainSessionStorage)

	userAuthHandler := webserver.NewUserAuthHandler(userAuthService)
	lobbyHandler := webserver.NewLobbyHandler(userAuthService)
	//webSocketHandler := handlers.NewWebSocketHandler(userAuthService, userWebSocketManageService)

	e := echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format:           "${time_custom} | ${method} ${uri} Status: ${status} Latency: ${latency} Bytes out: ${bytes_out}\n",
		CustomTimeFormat: "02.01.2006 15:04:05.000",
	}))
	e.Use(middleware.Recover())

	e.POST("/register", userAuthHandler.Register)
	e.POST("/authorize", userAuthHandler.Authorize)
	e.GET("/rooms", lobbyHandler.GetRoomsInfo)
	//e.GET("/ws", webSocketHandler.Handle)

	mlog.Fatal(e.Start(fmt.Sprintf("%s:%s", conf.Server.Host, conf.Server.Port)))
}

func runDBMigrate(dsn string, source *bindata.AssetSource) {
	d, err := bindata.WithInstance(source)
	if err != nil {
		mlog.Fatal(err)
	}

	m, err := migrate.NewWithSourceInstance("go-bindata", d, dsn)
	if err != nil {
		mlog.Fatal(err)
	}

	if err = m.Up(); err != nil {
		if err == migrate.ErrNoChange {
			mlog.Info("database migration: %s", err)
		} else {
			mlog.Fatal(err)
		}
	}
}
