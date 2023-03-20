package httpServer

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CorsSettings() *gin.HandlerFunc {
	c := cors.New(cors.Config{
		AllowOrigins: []string{
			"http://127.0.0.1:5173", "http://localhost:5173", "http://specialist-wbx.dl.wb.ru:8080",
			"http://localhost:8080", "http://127.0.0.1:8080"},
		AllowMethods: []string{
			"GET",
			"POST",
			"PUT",
			"PATCH",
			"DELETE",
			"OPTIONS"},
		AllowHeaders: []string{
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"X-CSRF-Token",
			"Authorization",
			"accept",
			"origin",
			"Cache-Control",
			"X-Requested-With",
			"Access-Control-Allow-Origin'"},
		ExposeHeaders: []string{
			"Content-Length"},
		AllowWildcard:    true,
		AllowCredentials: true,
	})
	return &c
}
