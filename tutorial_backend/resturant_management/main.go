package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

// var foodCollection *mongo.Collection = database.OpenCollection
func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "4800"
	}
	router := gin.New()

	// router.Use(gin.Logger())
	// routes.UseRouter(router)
	// router.Use(middleware.Authentication())
	router.Run(":" + port)
}
