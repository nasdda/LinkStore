package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/nasdda/linkstore/internal/interfaces"
	"github.com/nasdda/linkstore/pkg/logger"
)

type tagRoutes struct {
	tag interfaces.Tag
	l   logger.Logger
}

func newTagRoutes(handler *gin.RouterGroup, link interfaces.Link) {
	l := logger.NewLogger("[TAG]")
	r := &linkRoutes{link: link, l: l}

	h := handler.Group("/tag")
	{
		h.GET("/", r.getLink)
		h.POST("/", r.createLink)
	}
}
