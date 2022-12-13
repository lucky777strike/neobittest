package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAllTasks(c *gin.Context) {
	tasks, err := h.service.GetAllTasks()
	if err != nil {
		newErrorResponse(c, 404, fmt.Sprint(err))
		return
	}
	c.JSON(200, tasks)
}

func (h *Handler) RemoveAllTasks(c *gin.Context) {
	err := h.service.RemoveAllTasks()
	if err != nil {
		newErrorResponse(c, 500, fmt.Sprint(err))
		return
	}
	c.JSON(200, true)
}
