package main

import (
	database "github.com/binhkid2/gogin-surrealdb-start/Database"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	database.Connect()

	router.GET("/biolinks", database.GetAllBioLinks)
	router.GET("/biolink/:id", database.GetBioLink)
	router.POST("/biolinks", database.CreateBioLink)
	router.DELETE("/biolink/:id", database.DeleteBioLink)
	router.PUT("/biolink/:id", database.UpdateBioLink)
	// Close the database connection when the application exits
	defer database.Close()
	router.Run("localhost:8080")

}
