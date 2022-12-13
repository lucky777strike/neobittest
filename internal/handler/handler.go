package handler

import (
	"camparser/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func New(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.LoadHTMLGlob("./public/map/index.html")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	router.Static("/map", "./public/map")
	router.Static("/parse", "./public/parse")
	api := router.Group("/api")
	{
		api.GET("/test", h.test)
		api.GET("/getAllCams", h.getCamsMap)
		api.GET("/removeAllCam", h.RemoveAllCam)
		api.GET("/removeCam", h.RemoveCam)

		api.GET("/parse", h.ParseCams)
		api.GET("/getAllTasks", h.GetAllTasks)
		api.GET("/removeAllTasks", h.RemoveAllTasks)

	}
	return router
}
