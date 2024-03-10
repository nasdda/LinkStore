package v1

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/nasdda/linkstore/internal/interfaces"
)

type linkRoutes struct {
	link interfaces.Link
	l    *log.Logger
}

func newLinkRoutes(handler *gin.RouterGroup, link interfaces.Link) {
	l := &log.Logger{}
	l.SetPrefix("[LINK]")
	r := &linkRoutes{link: link, l: l}

	h := handler.Group("/translation")
	{
		h.GET("/", r.history)
		h.POST("/", r.doTranslate)
	}
}

// type historyResponse struct {
// 	History []entity.Translation `json:"history"`
// }

func (r *linkRoutes) history(c *gin.Context) {
}

func (r *linkRoutes) doTranslate(c *gin.Context) {
	// c.JSON(http.StatusOK, translation)
}
