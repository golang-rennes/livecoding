package main

import (
	"github.com/golang-rennes/livecoding/config"
	"github.com/golang-rennes/livecoding/models"
	"github.com/golang-rennes/livecoding/server"
)

func main() {
	c := &config.ServerConfig{
		DBHost:     "localhost",
		DBUser:     "postgres",
		DBPassword: "postgres",
		DBName:     "postgres",
		DBSSLMode:  "disable",
		HTTPPort:   8080,
		AdminList:  []string{"fsamin"},
	}

	models.DBConfig(c)

	e := server.New(c)
	e.Start()

}
