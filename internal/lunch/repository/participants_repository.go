package repository

import (
	"bc-labs-lunch-bot/ent"
	"sync"
)

type ParticipantsRepository interface {
	SaveParticipants(usersNames []string) ([]*ent.LunchParticipant, error)
}

var (
	onceParticipantsRepository sync.Once
	participantsRepository     ParticipantsRepository
)

func NewParticipantsRepository(client *ent.Client) ParticipantsRepository {
	if participantsRepository == nil {
		onceParticipantsRepository.Do(func() {
			participantsRepository = &participantsDB{
				client: client,
			}
		})
	}
	return participantsRepository
}
