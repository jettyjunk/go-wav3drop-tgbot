package main

import (
	"context"
	"go-wav3drop-tgbot/internal/bot"
	"go-wav3drop-tgbot/internal/config"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env", err)
	}

	cfg, err := config.Load()
	if err != nil {
		log.Fatal("failed to load config: ", err)
	}

	b, err := bot.NewBot(cfg.Token)
	if err != nil {
		log.Fatal("failed to create bot", err)
	}

	ctx := context.Background()
	log.Print("Wav3DropBot started")
	b.Start(ctx)
}
