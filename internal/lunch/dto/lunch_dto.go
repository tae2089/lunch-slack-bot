package dto

import "time"

type LunchDto struct {
	payer          string
	paymentDate    time.Time
	restaurantName string
	cafeName       string
}

func NewLunchDto(payer, restaurantName, cafeName string, paymentDate time.Time) *LunchDto {
	return &LunchDto{
		payer:          payer,
		paymentDate:    paymentDate,
		restaurantName: restaurantName,
		cafeName:       cafeName,
	}
}

func (l *LunchDto) Payer() string {
	return l.payer
}

func (l *LunchDto) PaymentDate() time.Time {
	return l.paymentDate
}

func (l *LunchDto) RestaurantName() string {
	return l.restaurantName
}

func (l *LunchDto) CafeName() string {
	return l.cafeName
}
