package main

import (
	"project/internal/application/services/links"
	linkhandlers "project/internal/controllers/link_handlers"
	"project/internal/domain/models"
	"project/internal/infrastructure"
	"strconv"

	"github.com/gin-gonic/gin"
)

// TODO Wire DI
func main() {
	router := gin.Default()

	linkRepo := infrastructure.NewLinksInMemoryRepo()
	linkService := links.NewLinkService(linkRepo)
	linkHandler := linkhandlers.NewLinkHandler(linkService)

	router.POST("/link", linkHandler.PostLink)
	router.GET("/link/:shortUrl", linkHandler.GetLink)

	router.Run(models.Domain + ":" + strconv.Itoa(models.Port))
}
