package config

import (
	"github.com/joho/godotenv"
	"github.com/tkanos/gonfig"
	"log"
)

type ServerConfig struct {
	Host HostConfig
	Database DBConfig
}

type HostConfig struct {
	ServerHost string `env:"SERVER_HOST"`
	ServerStartPort string `env:"SERVER_START_PORT"`
}

type DBConfig struct {
	Host   string `env:"DB_HOST"`
	Port       int    `env:"DB_PORT"`
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
	Name     string `env:"DB_NAME"`
}

func NewServerConfig() *ServerConfig {
	var err error

	var hostConfig HostConfig
	err = getConfig(&hostConfig)
	if err != nil {
		log.Println(err)
	}

	var dbConfig DBConfig
	err = getConfig(&dbConfig)
	if err != nil {
		log.Println(err)
	}

	return &ServerConfig{
		Host: hostConfig,
		Database: dbConfig,
	}
}

func getConfig(configType interface{}) (err error) {
	err = godotenv.Load(".env")
	if err != nil {
		log.Println(err)
	}
	return gonfig.GetConf("", configType)
}
