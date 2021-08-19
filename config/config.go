package config

import (
	"github.com/joho/godotenv"
	"github.com/tkanos/gonfig"
)

type ServerConfig struct {
	Server   Server
	Database DBConfig
	Token TokenConfig
}

type Server struct {
	Host string `env:"SERVER_HOST"`
	Port string `env:"SERVER_START_PORT"`
}

type DBConfig struct {
	Host   string `env:"DB_HOST"`
	Port       int    `env:"DB_PORT"`
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
	Name     string `env:"DB_NAME"`
}

type TokenConfig struct {
	TokenLength int `env:"TOKEN_LENGTH"`
	TokenSymbols string `env:"TOKEN_SYMBOLS"`
}

func NewServerConfig() *ServerConfig {
	var err error

	var hostConfig Server
	err = getConfig(&hostConfig)
	if err != nil {
		panic(err)
	}

	var dbConfig DBConfig
	err = getConfig(&dbConfig)
	if err != nil {
		panic(err)
	}

	var tokenConfig TokenConfig
	err = getConfig(&tokenConfig)
	if err != nil {
		panic(err)
	}

	return &ServerConfig{
		Server: hostConfig,
		Database:   dbConfig,
		Token:      tokenConfig,
	}
}

func getConfig(configType interface{}) (err error) {
	err = godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	return gonfig.GetConf("", configType)
}
