package app

import (
	appbot "github.com/dahaev/bottesting/internal/app/bot"
)

type Application struct {
	BotApp *appbot.Bot
}

func New() (*Application, error) {
	const token = "6859994276:AAFZz5JkEsZ_WbTs7Z1ZPNjBjRVte8DF6fg"
	//repo, _ := repository.New()
	//serviceInterface := service.New(repo)
	tgBot, err := appbot.New(token)
	if err != nil {
		return nil, err
	}
	return &Application{BotApp: tgBot}, err
}

func (a *Application) Start() {
	a.BotApp.Start()
}
