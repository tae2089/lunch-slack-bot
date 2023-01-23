package slack

import (
	"github.com/slack-go/slack"
	"net/http"
)

func SlashCommandParse(request http.Request) (string, error) {
	s, err := slack.SlashCommandParse(&request)
	if err != nil {
		return "", err
	}
	return s.TriggerID, nil
}
