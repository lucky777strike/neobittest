package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// @Summary getAllCams
// @Tags Cams
// @Description Get list of cameras
// @ID getAllCams
// @Accept  json
// @Produce  json
// @Success 200 {integer} bool true
// @Failure 404 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/getAllCams [get]
func (h *Handler) getCamsMap(c *gin.Context) {
	cams, err := h.service.GetAllCams()
	if err != nil {
		newErrorResponse(c, 404, fmt.Sprint(err))
		return
	}
	c.JSON(200, cams)
}

// @Summary removeCam
// @Tags Cams
// @Description Remove camera
// @ID removeCam
// @Accept  json
// @Produce  json
// @Param    id   path      int  true  "Camera ID"
// @Success 200 {integer} bool true
// @Failure 404 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/removeCam [get]
func (h *Handler) RemoveCam(c *gin.Context) {
	id := c.Query("id")
	err := h.service.RemoveCam(id)
	if err != nil {
		newErrorResponse(c, 404, fmt.Sprint(err))
		return
	}

	c.JSON(200, true)
}

// @Summary removeAllCam
// @Tags Cams
// @Description Remove all camers
// @ID removeAllCam
// @Accept  json
// @Produce  json
// @Success 200 {integer} bool true
// @Failure 404 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/removeAllCam [get]
func (h *Handler) RemoveAllCam(c *gin.Context) {
	err := h.service.RemoveAllCam()
	if err != nil {
		newErrorResponse(c, 500, fmt.Sprint(err))
		return
	}
	c.JSON(200, true)
}

// @Summary parse
// @Tags Parse
// @Description Parse shodan
// @ID parse
// @Accept  json
// @Produce  json
// @Param    query   path      string  true  "query"
// @Success 200 {integer} bool true
// @Failure 404 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/parse [get]
func (h *Handler) ParseCams(c *gin.Context) {
	query := c.Query("query")
	//log.Println(query)
	h.service.Parse(query)
	c.JSON(200, true)
}
