package main

import (
	"log"
	"project/internal/di"
	"project/internal/settings"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	conf := settings.NewConfig(".env")

	linkHandler, err := di.InitializeHandler(conf)
	if err != nil {
		log.Fatalf("failed to initialize dependencies: %v", err)
	}

	router := gin.Default()

	router.POST("/link", linkHandler.PostLink)
	router.GET("/link/:shortUrl", linkHandler.GetLink)

	router.Run(conf.Host + ":" + strconv.Itoa(conf.Port))
}
