package service

import (
	"bc-labs-lunch-bot/internal/lunch/dto"
	lunch "bc-labs-lunch-bot/internal/lunch/repository"
	local "bc-labs-lunch-bot/internal/slack"
	"fmt"
	"github.com/slack-go/slack"
	"log"
	"time"
)

var _ CallbackService = (*callbackServiceImpl)(nil)

type callbackServiceImpl struct {
	lunchRepository       lunch.LunchRepository
	participantRepository lunch.ParticipantsRepository
}

func (c callbackServiceImpl) SaveLunchPayment(i slack.InteractionCallback, token string) error {

	payerSelect := i.View.State.Values["payer selector"]["payer_select"]
	userSelect := i.View.State.Values["users selector"]["users_select"]
	restaurantField := i.View.State.Values["Restaurant Name"]["restaurantName"]
	cafeField := i.View.State.Values["Cafe Name"]["cafeName"]
	paymentDateField := i.View.State.Values["Payment Date"]["paymentDate"]
	paymentDate, err := time.Parse("20060102", paymentDateField.Value)
	if err != nil {
		return err
	}
	payerName, err := local.GetUserProfile(payerSelect.SelectedUser)
	if err != nil {
		return err
	}
	users, err := local.GetUsersRealName(userSelect.SelectedUsers...)
	if err != nil {
		return err
	}
	log.Println("users", users)
	participant, err := c.participantRepository.SaveParticipants(users)
	if err != nil {
		return err
	}

	lunchDto := dto.NewLunchDto(payerName, restaurantField.Value, cafeField.Value, paymentDate)

	_ = c.lunchRepository.SaveLunchPayment(*lunchDto, participant)
	msg := fmt.Sprintf("%s 님 점심 결제 등록이 정상적으로 완료됐습니다.!", payerName)

	err = local.PostMessage(payerSelect.SelectedUser,
		slack.MsgOptionText(msg, false),
		slack.MsgOptionAttachments(),
	)

	if err != nil {
		return err
	}

	return nil
}
