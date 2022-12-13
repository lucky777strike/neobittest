package handler

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) test(c *gin.Context) {
	id := c.Query("id")
	if id != "12345" {
		newErrorResponse(c, 404, "no id")
		return
	}
	c.String(200, "id setted as %s", id)
}
