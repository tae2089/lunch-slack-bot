package repository

import (
	"bc-labs-lunch-bot/ent"
	"context"
)

var _ ParticipantsRepository = (*participantsDB)(nil)

type participantsDB struct {
	client *ent.Client
}

func (p *participantsDB) SaveParticipants(usersNames []string) ([]*ent.LunchParticipant, error) {

	bulk := make([]*ent.LunchParticipantCreate, len(usersNames))
	for i, v := range usersNames {
		bulk[i] = p.client.LunchParticipant.Create().SetName(v)
	}
	savedLunchParticipantList, err := p.client.LunchParticipant.CreateBulk(bulk...).Save(context.Background())
	if err != nil {
		return nil, err
	}
	return savedLunchParticipantList, nil
}
