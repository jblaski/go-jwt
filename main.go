package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jblaski/go-jwt/controllers"
	"github.com/jblaski/go-jwt/database"
	"github.com/jblaski/go-jwt/middlewares"
	"github.com/jblaski/go-jwt/models"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	//here
	api := r.Group("/api")
	{
		public := api.Group("/public")
		{
			public.POST("/login", controllers.Login)
			public.POST("/signup", controllers.Signup)
		}

		protected := api.Group("protected").Use(middlewares.Autho())
		{
			protected.GET("/profile", controllers.Profile)
		}
	}

	return r
}

func main() {
	err := database.InitDatabase()
	if err != nil {
		log.Fatal("could not create database", err)
	}

	database.GlobalDB.AutoMigrate(&models.User{})

	r := setupRouter()
	r.Run(":8080")
}
