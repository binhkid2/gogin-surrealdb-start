package database

import (
	"fmt"
	"net/http"

	model "github.com/binhkid2/gogin-surrealdb-start/Model"
	"github.com/gin-gonic/gin"
	"github.com/surrealdb/surrealdb.go"
)

// getAlbums responds with the list of all albums as JSON.
func GetBioLinks(c *gin.Context) {
	var err error
	// access database use namespace and database
	if _, err = db.Use("test", "test"); err != nil {
		panic(err)
	}
	//Select all query
	biolinks, err := db.Select("biolinks")
	if err != nil {
		panic(err)
	}
	c.IndentedJSON(http.StatusOK, biolinks)

}

var (
	biolinks = map[string]model.BioLink{}
)

type TaskHandler struct {
	DB *surrealdb.DB
}

func CreateBioLink(c *gin.Context) {
	var err error
	// access database use namespace and database
	if _, err = db.Use("test", "test"); err != nil {
		panic(err)
	}
	biolink := new(model.BioLink)
	if err := c.ShouldBindJSON(biolink); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Error while binding json: %s", err.Error()),
		})
		return
	}

	db.Create("biolinks", biolink)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Error while creating task: %s", err.Error()),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"biolink": biolink,
	})
}
func DeleteBioLink(c *gin.Context) {
	var err error
	// access database use namespace and database
	if _, err = db.Use("test", "test"); err != nil {
		panic(err)
	}

	id := c.Param("id")

	if _, err = db.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Error while Delete task: %s", err.Error()),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"deleted": true})

}
func UpdateBioLink(c *gin.Context) {
	var err error
	// access database use namespace and database
	if _, err = db.Use("test", "test"); err != nil {
		panic(err)
	}
	//Update query
	id := c.Param("id")
	biolink := new(model.BioLink)
	if err := c.ShouldBindJSON(biolink); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Error while binding json: %s", err.Error()),
		})
		return
	}
	if _, err = db.Update(id, map[string]interface{}{
		"title":    biolink.Title,
		"link":     biolink.Link,
		"isPublic": biolink.IsPublic,
	}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Error while update task: %s", err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"updated": true,
	})
}
