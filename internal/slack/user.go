package slack

import (
	"github.com/slack-go/slack"
	"log"
)

func GetUserProfile(userId string) (string, error) {
	userProfile, err := client.GetUserProfile(
		&slack.GetUserProfileParameters{
			UserID: userId,
		},
	)
	if err != nil {
		log.Println(err)
		return "", err
	}
	log.Println(userProfile.RealName)
	log.Println(userProfile.DisplayName)
	return userProfile.RealName, nil
}

func GetUsersRealName(userId ...string) ([]string, error) {
	users, err := client.GetUsersInfo(userId...)
	userRealNameList := []string{}
	if err != nil {
		return userRealNameList, err
	}
	for _, user := range *users {
		userRealNameList = append(userRealNameList, user.RealName)
	}
	return userRealNameList, nil
}
