package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func loadRoutes() http.Handler {
	mux := gin.New()
	mux.Use(gin.Logger())
	
	mux.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "service online"})
	})

	return mux
}
