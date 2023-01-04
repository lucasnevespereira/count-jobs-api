package router

import (
	"count-jobs/api/handler"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	router := gin.New()

	router.GET("/", handler.Status())
	router.GET("/api", handler.Collect())

	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
	}))

	return router
}
