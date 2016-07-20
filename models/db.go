package models

import (
	"fmt"

	"github.com/golang-rennes/livecoding/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

//docker run --name postgres -e POSTGRES_PASSWORD=postgres -d -p 5432:5432 postgres

func DBConfig(c *config.ServerConfig) error {
	var err error
	connStr := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=%s password=%s", c.DBHost, c.DBUser, c.DBName, c.DBSSLMode, c.DBPassword)
	DB, err = gorm.Open("postgres", connStr)

	var user User
	DB.AutoMigrate(&user)

	return err
}
