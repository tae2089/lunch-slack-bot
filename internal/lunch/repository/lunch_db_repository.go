package repository

import (
	"bc-labs-lunch-bot/ent"
	"bc-labs-lunch-bot/internal/lunch/dto"
	"context"
	"log"
)

var _ LunchRepository = (*lunchDB)(nil)

type lunchDB struct {
	client *ent.Client
}

func (l *lunchDB) SaveLunchPayment(lunchDto dto.LunchDto, participants []*ent.LunchParticipant) error {

	log.Println("start SaveLunchPayment")
	_, err := l.client.Lunch.Create().
		SetPayer(lunchDto.Payer()).
		SetPaymentTime(lunchDto.PaymentDate()).
		SetCafeName(lunchDto.CafeName()).
		SetRestaurantName(lunchDto.RestaurantName()).
		AddParticipant(participants...).
		Save(context.Background())

	if err != nil {
		log.Println("error issue SaveLunchPayment")
		log.Println(err.Error())
		return err
	}
	log.Println("end SaveLunchPayment")
	return nil
}
