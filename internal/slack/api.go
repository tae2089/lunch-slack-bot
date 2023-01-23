package slack

import (
	"bc-labs-lunch-bot/internal/config"
	"fmt"
	"github.com/slack-go/slack"
	"log"
)

var (
	client *slack.Client
)

func init() {
	client = slack.New(config.GetAccessBotToken())
}

func OpenView(triggerId string, modalRequest slack.ModalViewRequest) error {
	_, err := client.OpenView(triggerId, modalRequest)
	if err != nil {
		log.Printf("Error opening view: %s", err)
		return err
	}
	return nil
}

func PostMessage(channelId string, options ...slack.MsgOption) error {
	_, _, err := client.PostMessage(channelId, options...)
	if err != nil {
		fmt.Printf(err.Error())
		return err
	}
	return nil
}
