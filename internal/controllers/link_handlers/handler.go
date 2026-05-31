package linkhandlers

import (
	"net/http"
	contracts "project/internal/application/contracts"
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
	var newLink contracts.LinkData

	if err := ctx.ShouldBindJSON(&newLink); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, status := handler.linkService.AddNewLink(newLink)

	switch status / 100 {
	case 2:
		ctx.JSON(status, result)
	default:
		ctx.JSON(status, gin.H{"error": http.StatusText(status)})
	}

}

func (handler *LinkHandler) GetLink(ctx *gin.Context) {
	link := ctx.Param("shortUrl")
	if len(link) == 0 {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}

	shortLink := contracts.ShortLinkData{ShortLink: link}

	result, status := handler.linkService.ExtractFullLink(shortLink)

	switch status / 100 {
	case 3:
		ctx.Redirect(status, result.FullUrl)
	default:
		ctx.JSON(status, gin.H{"error": http.StatusText(status)})
	}

}
