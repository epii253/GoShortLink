package linkhandlers

import (
	"net/http"
	"project/internal/application/services/links"

	"github.com/gin-gonic/gin"
)

type LinkHandler struct {
	linkService links.ILinkService
}

func NewLinkHandler(linkService links.ILinkService) *LinkHandler {
	return &LinkHandler{
		linkService: linkService}
}

func (handler *LinkHandler) PostLink(ctx *gin.Context) {
	var newLink links.LinkData

	if err := ctx.ShouldBindJSON(&newLink); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, status := handler.linkService.AddNewLink(newLink)

	switch status / 100 {
	case 4:
		ctx.JSON(status, result)
	default:
		ctx.JSON(http.StatusCreated, result)
	}

}

func (handler *LinkHandler) GetLink(ctx *gin.Context) {
	var shortLink links.ShortLinkData

	if err := ctx.ShouldBindJSON(&shortLink); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, status := handler.linkService.ExtractFullLink(shortLink)

	switch status / 100 {
	case 4:
		ctx.JSON(status, nil)
	default:
		ctx.Redirect(status, result.FullUrl)
	}

}
