package main

import (
	"project/internal/application/services/links"
	linkhandlers "project/internal/controllers/link_handlers"
	"project/internal/infrastructure"
	"project/internal/settings"
	"strconv"

	"github.com/gin-gonic/gin"
)

// TODO Wire DI
func main() {
	conf := settings.NewConfig("../../.env")

	router := gin.Default()

	linkRepo := infrastructure.NewLinksInMemoryRepo()
	linkService := links.NewLinkService(linkRepo)
	linkHandler := linkhandlers.NewLinkHandler(linkService)

	router.POST("/link", linkHandler.PostLink)
	router.GET("/link/:shortUrl", linkHandler.GetLink)

	router.Run(conf.Host + ":" + strconv.Itoa(conf.Port))
}
