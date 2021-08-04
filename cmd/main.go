package main

import (
	"backgammon/config"
	"backgammon/infrastructure/dal"
	"backgammon/infrastructure/dal/migrations"
	"backgammon/infrastructure/httphandlers"
	"backgammon/infrastructure/websockethandler"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	bindata "github.com/golang-migrate/migrate/v4/source/go_bindata"
	"github.com/labstack/echo/v4"
)

func main() {
	config := config.NewServerConfig()

	database := dal.NewDatabaseConnector(config)
	defer database.CloseDatabaseConnection()

	s:=bindata.Resource(migrations.AssetNames(), migrations.Asset)
	runDBMigrate(fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		config.Database.User,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.Name), s)

	webSocket := websockethandler.NewWebSocketHandler()
	userRegistrator := httphandlers.NewUserRegistrator()

	e := echo.New()
	e.GET("/", func(c echo.Context) error {return nil})
	e.GET("/ws", webSocket.Handle)
	e.POST("/register", userRegistrator.Handle)

	e.Logger.Fatal(e.Start(config.Host.ServerStartPort))
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
