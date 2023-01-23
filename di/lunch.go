package di

import (
	"bc-labs-lunch-bot/internal/config"
	"bc-labs-lunch-bot/internal/lunch/repository"
	"bc-labs-lunch-bot/internal/lunch/service"
)

func InitializeLunchServiceImpl() service.LunchService {
	lunchRepository := repository.NewLunchRepository(config.OpenDB())
	lunchService := service.NewLunchService(lunchRepository)
	return lunchService
}
