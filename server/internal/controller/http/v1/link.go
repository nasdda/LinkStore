package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nasdda/linkstore/internal/interfaces"
	"github.com/nasdda/linkstore/pkg/logger"
)

type linkRoutes struct {
	link interfaces.Link
	l    logger.Logger
}

func newLinkRoutes(handler *gin.RouterGroup, link interfaces.Link) {
	l := logger.NewLogger("[LINK]")
	r := &linkRoutes{link: link, l: l}

	h := handler.Group("/link")
	{
		h.GET("/", r.getLink)
		h.POST("/", r.createLink)
	}
}

// type historyResponse struct {
// 	History []entity.Translation `json:"history"`
// }

func (r *linkRoutes) getLink(c *gin.Context) {
	c.String(http.StatusOK, "Hello, World!")
}

func (r *linkRoutes) createLink(c *gin.Context) {
	// c.JSON(http.StatusOK, translation)
}
