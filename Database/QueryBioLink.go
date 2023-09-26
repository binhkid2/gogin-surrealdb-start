package database

import (
	"fmt"
	"net/http"

	model "github.com/binhkid2/gogin-surrealdb-start/Model"
	"github.com/gin-gonic/gin"
)

// getAlbums responds with the list of all albums as JSON.
func GetAllBioLinks(c *gin.Context) {
	var err error

	//Select all query
	biolinks, err := db.Select("biolinks")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve biolinks",
		})
		return
	}
	c.IndentedJSON(http.StatusOK, biolinks)
}

var (
	biolinks = map[string]model.BioLink{}
)

func GetBioLink(c *gin.Context) {
	var err error
	id := c.Param("id")
	//Select all query
	biolink, err := db.Select(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "No biolink have that id",
		})
		return
	}
	c.IndentedJSON(http.StatusOK, biolink)
}
func CreateBioLink(c *gin.Context) {
	var err error
	biolink := model.BioLink{}
	if err := c.ShouldBindJSON(&biolink); err != nil {
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
	//Update query
	id := c.Param("id")
	biolink := model.BioLink{}
	if err := c.ShouldBindJSON(&biolink); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Error while binding json: %s", err.Error()),
		})
		return
	}
	if _, err = db.Update(id, map[string]interface{}{
		"title":      biolink.Title,
		"link":       biolink.Link,
		"isPublic":   biolink.IsPublic,
		"created_at": biolink.CreatedAt,
		"updated_at": biolink.UpdatedAt,
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
