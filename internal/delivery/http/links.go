package http

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (h *Handler) LinksHandlerInit(api *gin.RouterGroup) {
	api.GET("/:link", h.GetLink)
	api.POST("/link", h.ShortLink)
}

func (h *Handler) GetLink(ctx *gin.Context) {
	var link = ctx.Param("link")
	shtLnk, err := h.services.Links.FindLink(link)
	if err != nil {
		ctx.JSON(http.StatusNotFound, "Link not found")
		log.Println(err.Error())
		return
	}

	ctx.Redirect(http.StatusMovedPermanently, shtLnk)
	return
}

func (h *Handler) ShortLink(ctx *gin.Context) {
	var link = ctx.Param("link")

	var shortLink string
	var err error

	if shortLink, err = h.services.Links.FindLink(link); err != nil {
		shortLink, err = h.services.Links.ShortLink(link)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}
	}

	ctx.JSON(http.StatusCreated, shortLink)
	return
}
