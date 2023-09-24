package main

import (
	database "github.com/binhkid2/gogin-surrealdb-start/Database"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/biolinks", database.GetBioLinks)
	router.POST("/biolinks", database.CreateBioLink)
	router.DELETE("/biolink/:id", database.DeleteBioLink)
	router.PUT("/biolink/:id", database.UpdateBioLink)
	database.Connect()
	router.Run("localhost:8080")
}
