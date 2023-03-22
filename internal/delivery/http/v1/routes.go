package v1

import (
	"IntegrationLab2/pkg/httpServer"
	"github.com/gin-gonic/gin"
)

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	c := httpServer.CorsSettings()
	router.Use(*c)

	apiV1 := router.Group("/api/v1")
	{
		apiV1.GET("/faculties", h.getFaculties)
		apiV1.GET("/groups", h.getGroups)
		apiV1.GET("/lessons", h.getLessons)
	}

	return router
}
