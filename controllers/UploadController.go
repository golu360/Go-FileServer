package controllers

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golu360/go-file-server/schemas"
)

func HandleFileUpload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "No File received",
		})
		return
	}
	if err := c.SaveUploadedFile(file, "data/"+file.Filename); err != nil {
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
	var directories []string
	fs, err := os.ReadDir("data")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, f := range fs {
		if f.Type().IsDir() {
			directories = append(directories, f.Name())
		}
	}
	c.IndentedJSON(http.StatusOK, gin.H{
		"keys": directories,
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
	if _, err := os.Stat("data/" + request.KeyName); !os.IsNotExist(err) {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	err := os.Mkdir("data/"+request.KeyName, 0755)
	if err != nil {
		log.Println("Error Creating Key")
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}
	var directories []string
	fs, err := os.ReadDir("data")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, f := range fs {
		if f.Type().IsDir() {
			directories = append(directories, f.Name())
		}
	}
	c.IndentedJSON(http.StatusOK, gin.H{
		"keys": directories,
	})
}
