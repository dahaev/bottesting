package appbot

import (
	"context"
	"fmt"
	"log"

	"github.com/dahaev/bottesting/pkg/models"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Bot struct {
	bot     *tgbotapi.BotAPI
	service service
}

func New(token string, service service) (*Bot, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal("cannot start bot", err)
		return nil, err
	}
	return &Bot{
		bot:     bot,
		service: service,
	}, nil
}

func (b *Bot) Start() {
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60
	updates, err := b.bot.GetUpdatesChan(updateConfig)
	if err != nil {
		fmt.Println(err)
	}
	// Обработка входящих сообщений
	res, err := b.service.GetLadyReviews(context.Background(), "@derzkaya")
	fmt.Println(res)
	fmt.Println(err)
	for update := range updates {
		if update.Message == nil {
			continue
		}
		if err != nil {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "К сожалению вы еще не имеете доступа к группам. Достаточно пройти регистрацию если Вы хотите делать публикации или просто запросить доступ для общения!")
			//keyboard := tgbotapi.NewInlineKeyboardMarkup(
			//	tgbotapi.NewInlineKeyboardRow(
			//		tgbotapi.NewInlineKeyboardButtonData("Регистрация", "registration"),
			//		tgbotapi.NewInlineKeyboardButtonData("Доступ", "auth"),
			//	),
			//)
			//keyboard = tgbotapi.New
			//msg.ReplyMarkup = keyboard
			//b.bot.Send(msg)
			buttonRegistration := tgbotapi.NewKeyboardButton("Регистрация")
			buttonAuth := tgbotapi.NewKeyboardButton("Доступ")
			row := tgbotapi.NewKeyboardButtonRow(buttonRegistration, buttonAuth)
			keyboard := tgbotapi.NewReplyKeyboard(row)
			msg.ReplyMarkup = keyboard
			b.bot.Send(msg)
			b.HandleMessage(update.Message)

		} else {

		}
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		//menu := b.CreateMenu()
		//msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hello")
		//msg.ReplyMarkup = menu
		//b.bot.Send(msg)

		//photo := b.CreateMessage(update)
		//btn1 := tgbotapi.NewInlineKeyboardButtonData("Информация", "btn1")
		//btn2 := tgbotapi.NewInlineKeyboardButtonData("Отзывы", "btn2")
		//row := tgbotapi.NewInlineKeyboardRow(btn1, btn2)
		//keyboard := tgbotapi.NewInlineKeyboardMarkup(row)

		// Устанавливаем клавиатуру в сообщение
		//photo.ReplyMarkup = keyboard
		//b.bot.Send(photo)

		// Пример ответа на приветствие
		if update.Message.IsCommand() {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Привет! Я бот на Golang.")
			_, err = b.bot.Send(msg)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

func (b *Bot) HandleMessage(message *tgbotapi.Message) {
	switch message.Text {
	case "Регистрация":
		msg := tgbotapi.NewMessage(message.Chat.ID, "регистрация начата")
		b.bot.Send(msg)
	case "Доступ":

	}
}

func (b *Bot) CreateMessage(update tgbotapi.Update) tgbotapi.PhotoConfig {
	photoUrl := "2.jpg"
	photo := tgbotapi.NewPhotoUpload(update.Message.Chat.ID, photoUrl)
	fmt.Println(photo)
	photo.Caption = b.CreateText()
	return photo
}

func (b *Bot) CreateText() string {
	name := "🇷🇺 Karolina"
	ordens := "🏆 best blowjob"
	location := "Мытищи"
	money := "💶 от 4500"
	time := "🕞 до 03"
	account := "💌 @karolina"

	oby := "📋 ти английские тексты для начинающих предоставляют возможность бесплатно попрактиковаться в чтении и понимании онлайн. " +
		"Занятия по пониманию письменного английского языка расширят ваш словарный запас и улучшат понимание грамматики и порядка слов." +
		" Тексты ниже призваны помочь вам развиваться и предоставляют вам мгновенную оценку вашего прогресса.\n"

	rating := "⭐ 7.6 (10)"
	sendMessage := fmt.Sprintf("%s\n📍 %s\n%s\n%s\n%s\n%s\n%s\n%s", name, location, oby, ordens, money, time, rating, account)

	return sendMessage
}

func (b *Bot) CreateMenu() *tgbotapi.ReplyKeyboardMarkup {
	btn1 := tgbotapi.NewKeyboardButton("Registration")
	btn2 := tgbotapi.NewKeyboardButton("SomeThingElse")
	btn3 := tgbotapi.NewKeyboardButton("Row2")
	btn4 := tgbotapi.NewKeyboardButton("Row3")
	row1 := tgbotapi.NewKeyboardButtonRow(btn1, btn2)
	row2 := tgbotapi.NewKeyboardButtonRow(btn3, btn4)
	keyboard := tgbotapi.NewReplyKeyboard(row1, row2)
	return &keyboard

}

type service interface {
	CreateLadyAccount(ctx context.Context, account *models.Account) error
	GetAccountLady(ctx context.Context, userName string) (*models.Account, error)
	CreateDonAccount(ctx context.Context, username string) error
	GetDonAccount(ctx context.Context, userName string) (*models.DonAccount, error)
	CreateReview(ctx context.Context, ladyName string, donName string, review string, rating int) error
	GetLadyReviews(ctx context.Context, ladyName string) ([]models.Review, error)
}
