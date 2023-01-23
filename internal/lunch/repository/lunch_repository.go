package repository

import (
	"bc-labs-lunch-bot/ent"
	"bc-labs-lunch-bot/internal/lunch/dto"
	"sync"
)

type LunchRepository interface {
	SaveLunchPayment(lunchDto dto.LunchDto, participants []*ent.LunchParticipant) error
}

var (
	onceLunchRepository sync.Once
	lunchRepository     LunchRepository
)

func NewLunchRepository(client *ent.Client) LunchRepository {
	if lunchRepository == nil {
		onceLunchRepository.Do(
			func() {
				lunchRepository = &lunchDB{
					client: client,
				}
			})
	}
	return lunchRepository
}
