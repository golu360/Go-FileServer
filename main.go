package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golu360/go-file-server/controllers"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	var hostUrl = fmt.Sprintf("localhost:%s", viper.Get("PORT"))
	router := gin.Default()
	log.Println("Running on " + hostUrl)
	router.StaticFS("/data", http.Dir("data"))
	router.POST("/upload", controllers.HandleFileUpload)
	router.GET("/keys", controllers.GetFS)
	router.POST("/keys", controllers.CreateKey)
	router.Run(hostUrl)
}
