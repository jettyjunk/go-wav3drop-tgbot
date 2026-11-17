package config

import (
	"fmt"
	"os"
)

type Config struct {
	Token string
}

func Load() (Config, error) {
	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	if token == "" {
		return Config{}, fmt.Errorf("token tgbot not found")
	}

	return Config{
		Token: token,
	}, nil
}
