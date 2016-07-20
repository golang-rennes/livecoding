package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-rennes/livecoding/config"
	"github.com/golang-rennes/livecoding/models"
)

func GetUsersHandler(config *config.ServerConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		//users := []models.User{models.User{Username: "fsamin", Fullname: "François SAMIN", Email: "francois.samin@gmail.com"}}
		users := []models.User{}
		models.DB.Find(&users)

		c.JSON(http.StatusOK, users)
		c.Next()
	}
}

func PostNewUserHandler(config *config.ServerConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		var username string

		username = c.Query("username")
		if username == "" {
			username = c.PostForm("username")
		}
		if username == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "username is mandatory"})
			return
		}

		user := models.User{Username: username}

		for _, admin := range config.AdminList {
			if username == admin {
				user.IsAdmin = true
			}
		}

		if errs := models.DB.Create(&user).GetErrors(); len(errs) > 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"errors": errs})
			return
		}

		c.JSON(http.StatusCreated, user)
		c.Next()
	}
}

func GetUserHandler(config *config.ServerConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.Param("username")
		if username == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "username is mandatory"})
			return
		}
		var user models.User
		if errs := models.DB.Where("username = ?", username).First(&user).GetErrors(); len(errs) > 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"errors": errs})
			return
		}

		c.JSON(http.StatusOK, user)
		c.Next()
	}
}

func EditUserHandler(config *config.ServerConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.Param("username")
		if username == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "username is mandatory"})
			return
		}
		var user models.User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if username != user.Username {
			c.JSON(http.StatusBadRequest, gin.H{"error": "formdata doesn't match"})
			return
		}
		if errs := models.DB.Save(&user).GetErrors(); len(errs) > 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"errors": errs})
			return
		}
		c.JSON(http.StatusOK, user)
	}
}
