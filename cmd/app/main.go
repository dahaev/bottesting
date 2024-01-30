package main

import (
	"log"

	appbot "github.com/dahaev/bottesting/internal/app/bot"
)

func main() {
	const token = "6859994276:AAFZz5JkEsZ_WbTs7Z1ZPNjBjRVte8DF6fg"
	bot, err := appbot.New(token)
	if err != nil {
		log.Fatal(err)
	}
	bot.Start()
}
