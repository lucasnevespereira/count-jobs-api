package handler

import (
	"count-jobs/collector"
	"github.com/gin-gonic/gin"
)

func Status() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "up"})
	}
}

func Collect() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		resultData := collector.StartCollector(ctx.Query("term"), ctx.Query("location"), ctx.Query("country"))
		ctx.JSON(200, gin.H{"result": string(resultData)})
	}
}
