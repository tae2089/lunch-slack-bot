package config

import (
	"fmt"
	"os"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

var (
	slackbotConfig = &SlackBotConfig{}
	dbConfig       = &DbConfig{}
)

func init() {
	profile := os.Getenv("GO_PROFILE")
	if profile == "" || profile == "local" {
		godotenv.Load("./configs/.env")
	}
	if err := env.Parse(slackbotConfig); err != nil {
		panic(err)
	}
	if err := env.Parse(dbConfig); err != nil {
		panic(err)
	}
}

func GetAccessBotToken() string {
	return slackbotConfig.AccessToken
}

func GetSecretBotToken() string {
	return slackbotConfig.SecretToken
}

func GetDbUrl() string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s",
		dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Database)
}
