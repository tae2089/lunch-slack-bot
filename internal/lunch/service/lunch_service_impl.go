package service

import (
	"bc-labs-lunch-bot/internal/lunch/repository"
	"bc-labs-lunch-bot/internal/slack"
	"fmt"
	"net/http"
)

var _ LunchService = (*lunchServiceImpl)(nil)

type lunchServiceImpl struct {
	lunchRepository repository.LunchRepository
}

// RegisterLunchPayment implements LunchService
func (*lunchServiceImpl) RegisterLunchPayment(request http.Request) error {

	triggerId, err := slack.SlashCommandParse(request)
	if err != nil {
		return err
	}
	modalRequest := slack.GenerateModalRequest()
	err = slack.OpenView(triggerId, modalRequest)
	if err != nil {
		fmt.Printf("Error opening view: %s", err)
		return err
	}
	return nil
}
