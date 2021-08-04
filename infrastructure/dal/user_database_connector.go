package dal

import (
	"backgammon/config"
	"fmt"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

type DatabaseConnector struct {
	Database *sqlx.DB
}

func NewDatabaseConnector(sConfig *config.ServerConfig) *DatabaseConnector {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		sConfig.Database.Host, sConfig.Database.Port, sConfig.Database.User,
		sConfig.Database.Password, sConfig.Database.Name)

	db, err := sqlx.Connect("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return &DatabaseConnector{
		Database: db,
	}
}

func (d *DatabaseConnector) CloseDatabaseConnection() {
	err := d.Database.Close()
	if err != nil {
		log.Println(err)
	}
}