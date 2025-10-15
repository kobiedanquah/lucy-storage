package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *ServerApplication) loadRoutes() http.Handler {
	mux := gin.New()
	mux.Use(gin.Logger())
	mux.Use(gin.Recovery())

	mux.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "service online"})
	})

	mux.POST("/users", app.handler.CreateUser)
	mux.GET("/users/:id", app.handler.GetUser)
	mux.PATCH("/users/:id", app.handler.UpdateUserData)
	mux.DELETE("/users/:id", app.handler.DeleteUser)

	return mux
}
