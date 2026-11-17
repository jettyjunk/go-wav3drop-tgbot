package bot

import (
	"context"
	"log"

	telegram "github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func StartHandler(ctx context.Context, b *telegram.Bot, update *models.Update) {
	if update == nil || update.Message == nil {
		return
	}

	_, err := b.SendMessage(ctx, &telegram.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Hello!I Wab3DropBot",
	})
	if err != nil {
		log.Printf("failed to send message: %v", err)
	}
}
