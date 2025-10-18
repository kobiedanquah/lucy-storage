package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/primekobie/lucy/internal/handlers"
)

func (app *ServerApplication) loadRoutes() http.Handler {
	mux := gin.New()
	mux.Use(gin.Logger())
	mux.Use(gin.Recovery())

	mux.Use(cors.New(cors.Config{
		AllowMethods: []string{"POST", "GET", "DELETE", "PATCH", "OPTIONS"},
		AllowOrigins: []string{"http://localhost:5173"},
		AllowHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
	}))

	open := mux.Group("/api/v1")
	open.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"status":  "200",
			"message": "online",
		})
	})

	// users
	open.POST("/users/register", app.handler.CreateUser)
	open.POST("/users/login", app.handler.LoginUser)
	open.POST("/users/access", app.handler.GetUserAccessToken)
	open.POST("/users/verify", app.handler.VerifyUser)
	open.POST("/users/verify/request", app.handler.RequestVerification)

	protected := open.Group("/")
	protected.Use(handlers.Authentication())
	{
		//users
		protected.GET("/users/:id", app.handler.GetUser)
		protected.PATCH("/users/profile", app.handler.UpdateUserData)
		protected.DELETE("/users/:id", app.handler.DeleteUser)
	}

	return mux
}
