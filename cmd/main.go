package main

import (
	"backgammon/config"
	"backgammon/infrastructure/dal"
	"backgammon/infrastructure/dal/migrations"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	bindata "github.com/golang-migrate/migrate/v4/source/go_bindata"
)

func main() {
	fmt.Println("Hello backgammon!")

	config := config.NewServerConfig()

	fmt.Println(config)

	database := dal.NewDatabaseConnector(config)

	defer database.CloseDatabaseConnection()

	s:=bindata.Resource(migrations.AssetNames(), migrations.Asset)
	runDBMigrate(fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		config.Database.User,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.Name), s)
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
