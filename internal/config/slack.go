package config

type SlackBotConfig struct {
	AccessToken string `env:"SLACK_BOT_ACCESS_TOKEN"`
	SecretToken string `env:"SLACK_BOT_SECRET_TOKEN"`
}
