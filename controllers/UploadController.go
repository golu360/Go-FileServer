package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golu360/go-file-server/schemas"
	"github.com/golu360/go-file-server/utils"
)

func HandleFileUpload(c *gin.Context) {
	c.Request.ParseForm()

	var keyName string = c.PostForm("key")
	var fileUtils utils.FileUtils = utils.FileUtils{DirName: "data"}
	log.Println(fileUtils.KeyExists(keyName))
	if !fileUtils.KeyExists(keyName) {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "Invalid Key",
		})
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "No File received",
		})
		return
	}
	if err := c.SaveUploadedFile(file, "data/"+keyName+"/"+file.Filename); err != nil {
		log.SetPrefix("Upload Controller ")
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "File Upload Success",
	})

}

func GetFS(c *gin.Context) {
	var fileUtils utils.FileUtils = utils.FileUtils{DirName: "data"}

	c.IndentedJSON(http.StatusOK, gin.H{
		"keys": fileUtils.GetKeys(),
	})
}

func CreateKey(c *gin.Context) {
	var request schemas.CreateKeyRequest
	if err := c.BindJSON(&request); err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}
	var fileUtils utils.FileUtils = utils.FileUtils{DirName: "data"}
	var created bool = fileUtils.CreateKey(request.KeyName)
	if !created {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Could not create key",
		})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{
		"keys": fileUtils.GetKeys(),
	})
}
