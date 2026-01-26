package main

import (
	"project/internal/application/services/links"
	linkhandlers "project/internal/controllers/link_handlers"
	"project/internal/infrastructure"

	"github.com/gin-gonic/gin"
)

// TODO Wire DI
func main() {
	router := gin.Default()

	linkRepo := infrastructure.NewLinksInMemoryRepo()
	linkService := links.NewLinkService(linkRepo)
	linkHandler := linkhandlers.NewLinkHandler(linkService)

	router.POST("/link", linkHandler.PostLink)
}
