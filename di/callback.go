package di

import (
	"bc-labs-lunch-bot/internal/callback/service"
	"bc-labs-lunch-bot/internal/config"
	lunch "bc-labs-lunch-bot/internal/lunch/repository"
)

//CallBackService Dependency Injection
func InitializeCallbackServiceImpl() service.CallbackService {
	lunchRepository := lunch.NewLunchRepository(config.OpenDB())
	participantRepository := lunch.NewParticipantsRepository(config.OpenDB())
	callBackService := service.NewCallbackService(lunchRepository, participantRepository)
	return callBackService
}
