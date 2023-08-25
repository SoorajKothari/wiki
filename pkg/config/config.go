package config

import (
	"fmt"
	"gorm.io/gorm"
	"log"
	"os"
	"wiki/pkg/db"
)

type Config struct {
	DbHost     string `yaml:"dbHost"`
	DbName     string `yaml:"dbName"`
	DbPassword string `yaml:"dbPassword"`
	DbUser     string `yaml:"dbUser"`
	DB         *gorm.DB
	ServerPort string `yaml:"serverPort"`
}

var (
	dbHost     = "localhost:4000"
	dbName     = "books_db"
	dbPass     = "dummy"
	dbUser     = "root"
	serverPort = "4001"
)

func GetConfig() (*Config, error) {
	config := &Config{
		DbHost:     dbHost,
		DbName:     dbName,
		DbPassword: dbPass,
		DbUser:     dbUser,
		ServerPort: serverPort,
	}

	config.UpdateSecrets()
	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable", config.DbHost, config.DbUser, config.DbPassword, config.DbName)
	log.Println("DSN: " + connStr)
	log.Println("Initializing Database")
	config.DB = db.InitializeDB(connStr)

	return config, nil
}

func (config *Config) UpdateSecrets() {
	log.Println("Updating Secrets.")

	if secret, found := os.LookupEnv("DBHOST"); found {
		config.DbHost = secret
	}
	if secret, found := os.LookupEnv("DBNAME"); found {
		config.DbName = secret
	}
	if secret, found := os.LookupEnv("DBPASS"); found {
		config.DbPassword = secret
	}
	if secret, found := os.LookupEnv("DBUSER"); found {
		config.DbUser = secret
	}
	if secret, found := os.LookupEnv("SERVER_PORT"); found {
		config.ServerPort = secret
	}
	log.Println("Secrets Updated.")
}
