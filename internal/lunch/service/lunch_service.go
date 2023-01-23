package service

import (
	"bc-labs-lunch-bot/internal/lunch/repository"
	"net/http"
	"sync"
)

type LunchService interface {
	RegisterLunchPayment(http.Request) error
}

var (
	onceLunchService sync.Once
	lunchService     LunchService
)

func NewLunchService(lunchRepository repository.LunchRepository) LunchService {
	if lunchService == nil {
		onceLunchService.Do(
			func() {
				lunchService = &lunchServiceImpl{
					lunchRepository: lunchRepository,
				}
			})
	}
	return lunchService
}
