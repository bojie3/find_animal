package main

import (
	"net/http"

	"github.com/bojie/animal/backend/db"
	"github.com/bojie/animal/backend/user"
	"github.com/bojie/animal/backend/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	db.LoadEnvVar()
	db.SetupDatabase()
}

func main() {
	router := gin.Default()

	router.Use(middleware.CORSmiddleware())

	user.UserRoutes(router)

	router.NoRoute(func(ctx *gin.Context) { ctx.JSON(http.StatusNotFound, gin.H{"error": "no such api"}) })

	// Start listening and serving requests
	router.Run(":8080")

	defer db.DB.Close()
}
