package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// @Summary getAllTasks
// @Tags Tasks
// @ID getAllTasks
// @Accept  json
// @Produce  json
// @Success 200 {integer} bool true
// @Failure 404 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/getAllTasks [get]
func (h *Handler) GetAllTasks(c *gin.Context) {
	tasks, err := h.service.GetAllTasks()
	if err != nil {
		newErrorResponse(c, 404, fmt.Sprint(err))
		return
	}
	c.JSON(200, tasks)
}

// @Summary removeAllTasks
// @Tags Tasks
// @Description Remove All Tasks
// @ID removeAllTasks
// @Accept  json
// @Produce  json
// @Success 200 {integer} bool true
// @Failure 404 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/removeAllTasks [get]
func (h *Handler) RemoveAllTasks(c *gin.Context) {
	err := h.service.RemoveAllTasks()
	if err != nil {
		newErrorResponse(c, 500, fmt.Sprint(err))
		return
	}
	c.JSON(200, true)
}
