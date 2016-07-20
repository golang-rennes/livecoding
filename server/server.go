package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/golang-rennes/livecoding/config"
)

//Engine is the server struct
type Engine struct {
	*gin.Engine
	port   int
	config config.ServerConfig
}

//New intializes server from config
func New(c *config.ServerConfig) *Engine {
	var router = &Engine{gin.Default(), 3000, *c}
	router.routesConfig(c)
	return router
}

//Start starts the server
func (e *Engine) Start() {
	e.Run(fmt.Sprintf(":%d", e.config.HTTPPort))
}
