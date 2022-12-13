package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getCamsMap(c *gin.Context) {
	cams, err := h.service.GetAllCams()
	if err != nil {
		newErrorResponse(c, 404, fmt.Sprint(err))
		return
	}
	c.JSON(200, cams)
}

func (h *Handler) RemoveCam(c *gin.Context) {
	id := c.Query("id")
	err := h.service.RemoveCam(id)
	if err != nil {
		newErrorResponse(c, 404, fmt.Sprint(err))
		return
	}

	c.JSON(200, true)
}

func (h *Handler) RemoveAllCam(c *gin.Context) {
	err := h.service.RemoveAllCam()
	if err != nil {
		newErrorResponse(c, 500, fmt.Sprint(err))
		return
	}
	c.JSON(200, true)
}

func (h *Handler) ParseCams(c *gin.Context) {
	query := c.Query("query")
	//log.Println(query)
	h.service.Parse(query)
	c.JSON(200, true)
}
