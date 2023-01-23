package service

import (
	lunch "bc-labs-lunch-bot/internal/lunch/repository"
	"github.com/slack-go/slack"
	"sync"
)

type CallbackService interface {
	SaveLunchPayment(i slack.InteractionCallback, token string) error
}

var (
	onceCallbackService sync.Once
	callbackService     CallbackService
)

func NewCallbackService(lunchRepository lunch.LunchRepository, participantRepository lunch.ParticipantsRepository) CallbackService {
	if callbackService == nil {
		onceCallbackService.Do(
			func() {
				callbackService = &callbackServiceImpl{
					lunchRepository:       lunchRepository,
					participantRepository: participantRepository,
				}
			})
	}
	return callbackService
}
