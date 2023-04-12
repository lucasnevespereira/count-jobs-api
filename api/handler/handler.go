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
		c := collector.NewCollector(collector.NewIndeedSource(), collector.NewLinkedInSource())
		jobChan, err := c.Start(ctx, ctx.Query("term"), ctx.Query("location"))
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}

		var jobs []collector.Job
		for job := range jobChan {
			jobs = append(jobs, job)
		}

		ctx.JSON(200, gin.H{"jobs": jobs})
	}
}
