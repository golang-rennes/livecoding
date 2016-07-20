package server

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-rennes/livecoding/config"
	"github.com/golang-rennes/livecoding/controllers"
)

func (e *Engine) routesConfig(config *config.ServerConfig) {

	e.Any("/", func(c *gin.Context) {
		c.String(200, "welcome")
	})

	//On passe la structure de configuration à l'initialisation des handlers
	//Cela permet

	//$ curl -XGET localhost:8080/users
	e.GET("/users" /*basicAuth(),*/, controllers.GetUsersHandler(config))
	//$ curl -XPOST localhost:8080/users/new?username=fsamin
	e.POST("/users/new", controllers.PostNewUserHandler(config))
	//$ curl -XGET localhost:8080/users/fsamin
	e.GET("/users/:username", controllers.GetUserHandler(config))
	//$ curl -XPUT --data '{"ID":1,"Username":"fsamin","Fullname":"François SAMIN","Email":"francois.samin@gmail.com"}'  localhost:8080/users/fsamin
	e.PUT("/users/:username", controllers.EditUserHandler(config))

	/*
		users := e.Group("/users")
		{
			users.GET("", controllers.GetUsersHandler(config))
			users.POST("/new", controllers.PostNewUserHandler(config))
			users.GET("/:username", controllers.GetUserHandler(config))
			users.PUT("/:username", controllers.EditUserHandler(config))
		}
	*/

}
