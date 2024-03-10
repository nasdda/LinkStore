package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/nasdda/linkstore/internal/interfaces"
)

func NewRouter(handler *gin.Engine, link interfaces.Link) {
	// Options
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	// Routers
	h := handler.Group("/v1")
	{
		newLinkRoutes(h, link)
	}
}
