package main

import (
	"bc-labs-lunch-bot/internal/callback"
	"bc-labs-lunch-bot/internal/health"
	"bc-labs-lunch-bot/internal/lunch"
	"bc-labs-lunch-bot/internal/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	//health check group
	healthGroup := app.Group("/bc-lab")
	healthGroup.GET("/healthz", health.HealthCheck)

	//slack group
	slack := app.Group("/slack")
	slack.Use(middleware.VerifySlack())
	slack.POST("/slash", lunch.CreateModalLunchCosts)
	slack.POST("/call-back", callback.CallBack)
	app.Run(":3000")
}
