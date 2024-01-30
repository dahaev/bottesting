package app

import appbot "github.com/dahaev/bottesting/internal/app/bot"

type Application struct {
	BotApp *appbot.Bot
}

func (a *Application) Start() {
	a.BotApp.Start()
}
