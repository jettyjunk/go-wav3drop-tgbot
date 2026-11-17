package bot

import (
	telegram "github.com/go-telegram/bot"
)

func NewBot(token string) (*telegram.Bot, error) {
	b, err := telegram.New(token)
	if err != nil {
		return nil, err
	}

	b.RegisterHandler(
		telegram.HandlerTypeMessageText,
		"/start",
		telegram.MatchTypeExact,
		StartHandler,
	)

	return b, nil
}
