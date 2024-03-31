package appbot

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/dahaev/bottesting/pkg/models"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	bot         *tgbotapi.BotAPI
	service     service
	userState   map[int64]string
	userChatIDS map[int64]struct{}
}

type MessageData struct {
	MessageID int
	ChatID    int64
	Keyboard  *tgbotapi.InlineKeyboardMarkup
}

func New(token string) (*Bot, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal("cannot start bot", err)
		return nil, err
	}
	return &Bot{
		bot:         bot,
		userState:   make(map[int64]string),
		userChatIDS: make(map[int64]struct{}),
	}, nil
}

func (b *Bot) SendMainMenu(update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Для доступа к группе, необходимо пройти регистрацию или получить доступ")
	buttonRegistration := tgbotapi.NewKeyboardButton("Регистрация")
	buttonAuth := tgbotapi.NewKeyboardButton("Доступ")
	row := tgbotapi.NewKeyboardButtonRow(buttonRegistration, buttonAuth)
	keyboard := tgbotapi.NewReplyKeyboard(row)
	msg.ReplyMarkup = keyboard
	b.bot.Send(msg)
	b.userState[update.Message.Chat.ID] = "Registration/access"
}

func (b *Bot) SendAnnounce() {

	for key := range b.userChatIDS {
		btn1 := tgbotapi.NewInlineKeyboardButtonData("Подробнее", "btn1")
		btn2 := tgbotapi.NewInlineKeyboardButtonData("Контакты", "btn2")
		btn3 := tgbotapi.NewInlineKeyboardButtonData("О компании", "btn3")
		row1 := tgbotapi.NewInlineKeyboardRow(btn1, btn2)
		row2 := tgbotapi.NewInlineKeyboardRow(btn3)

		photo := b.CreateMessage(key)
		photo.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(row1, row2)
		_, err := b.bot.Send(photo)

		if err != nil {
			log.Println("huy")
		}
	}
}

func (b *Bot) Start() {
	b.bot.Debug = false

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := b.bot.GetUpdatesChan(u)
	ticker := time.NewTicker(20 * time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				b.SendAnnounce()
			}
		}
	}()

	for update := range updates {
		if update.Message != nil { // Если мы получили сообщение
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Добро пожаловать. Вы будете получать различные предложение по работе.")
			if _, err := b.bot.Send(msg); err != nil {
				log.Panic(err)
			}
			b.userChatIDS[update.Message.Chat.ID] = struct{}{}
			fmt.Println("WE ARE HEREEEE 99")
		}
		if update.CallbackQuery != nil {
			data := update.CallbackQuery.Data

			// Отвечаем пользователю

			switch data {
			case "btn2":
				chatID := update.CallbackQuery.Message.Chat.ID
				contacts := fmt.Sprintf("\n%s\n☎︎ - %s\n📧-%s", "Karolina", "+7 (267) 632-32-32", "@karolina")
				//msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, contacts)
				msg := tgbotapi.NewMessage(chatID, contacts)
				msg.ReplyToMessageID = update.CallbackQuery.Message.MessageID
				_, err := b.bot.Send(msg)
				if err != nil {
					log.Println(err)
				}
			case "btn3":
				chatID := update.CallbackQuery.Message.Chat.ID
				photo := tgbotapi.NewPhoto(chatID, tgbotapi.FilePath("ozon_logo.jpeg"))
				contacts := fmt.Sprintf("Ozon — мультикатегорийный маркетплейс, с развитой логистической сетью. Компания входит в топ-5 крупнейших российских ритейлеров и развивает бизнес в России, странах СНГ, Китае и Турции. По итогам первой половины 2023 года оборот маркетплейса превысил 675 млрд рублей. На Ozon торгует более 300 000 продавцов, которые предлагают на площадке более 200 млн товарных наименований. Платформу ежемесячно посещает более 65 млн пользователей Рунета, а число ее активных покупателей достигло 40 млн человек.")
				//msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, contacts)
				photo.Caption = contacts
				photo.ReplyToMessageID = update.CallbackQuery.Message.MessageID
				_, err := b.bot.Send(photo)
				if err != nil {
					log.Println(err)
				}
			}
		}
	}
}

//for update := range updates {
//	if update.Message == nil {
//		continue
//	}
//	if currentState, ok := b.userState[update.Message.Chat.ID]; ok {
//		if currentState == "Registration/access" {
//			if update.Message.Text == "Регистрация" {
//				//msgTG := tgbotapi.NewMessage(update.Message.Chat.ID, b.CreateText())
//				b.bot.Send(b.CreateMessage(update))
//			} else if update.Message.Text == "Доступ" {
//				msgTG := tgbotapi.NewMessage(update.Message.Chat.ID, "Вы нажали зарегестрироваться как Мудила")
//				b.bot.Send(msgTG)
//			}
//			regMsg := tgbotapi.NewMessage(update.Message.Chat.ID, "Давайте проведем регистрацию - это всего несколько простых шагов!")
//			buttonNext := tgbotapi.NewKeyboardButton("next")
//			rowReg := tgbotapi.NewKeyboardButtonRow(buttonNext)
//			keyboardReg := tgbotapi.NewReplyKeyboard(rowReg)
//			regMsg.ReplyMarkup = keyboardReg
//			b.bot.Send(regMsg)
//		}
//	} else {
//		b.SendMainMenu(update)
//	}
//}

//keyboard := tgbotapi.NewInlineKeyboardMarkup(
//	tgbotapi.NewInlineKeyboardRow(
//		tgbotapi.NewInlineKeyboardButtonData("Регистрация", "registration"),
//		tgbotapi.NewInlineKeyboardButtonData("Доступ", "auth"),
//	),
//)
//keyboard = tgbotapi.New
//msg.ReplyMarkup = keyboard
//b.bot.Send(msg)
//menu := b.CreateMenu()
//msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hello")
//msg.ReplyMarkup = menu
//b.bot.Send(msg)
//
//photo := b.CreateMessage(update)
//btn1 := tgbotapi.NewInlineKeyboardButtonData("Информация", "btn1")
//btn2 := tgbotapi.NewInlineKeyboardButtonData("Отзывы", "btn2")
//row := tgbotapi.NewInlineKeyboardRow(btn1, btn2)
//keyboard := tgbotapi.NewInlineKeyboardMarkup(row)
//
//Устанавливаем клавиатуру в сообщение
//photo.ReplyMarkup = keyboard
//b.bot.Send(photo)

// Пример ответа на приветствие
//
//	func (b *Bot) HandleMessage(message *tgbotapi.Message, update tgbotapi.Update) {
//		switch message.Text {
//		case "Регистрация":
//			msg := tgbotapi.NewMessage(message.Chat.ID, "регистрация начата")
//			b.bot.Send(msg)
//		case "Доступ":
//
//		}
//	}
func (b *Bot) CreateMessage(chatID int64) tgbotapi.PhotoConfig {
	//photofile, err := os.Open("1.jpg")
	//fmt.Println(err)
	photo := tgbotapi.NewPhoto(chatID, tgbotapi.FilePath("ozon.jpeg"))
	photo.Caption = b.CreateTextOzon()
	return photo
}

func (b *Bot) CreateTextOzon() string {
	//сompanyName := "OZON"
	location := "Московская область"
	salary := "💶 от 55 000"
	time := "🕞 смена до 10 часов"
	account := "💌 @karolina"

	oby := "📋 ти английские тексты для начинающих предоставляют возможность бесплатно попрактиковаться в чтении и понимании онлайн. " +
		"Занятия по пониманию письменного английского языка расширят ваш словарный запас и улучшат понимание грамматики и порядка слов." +
		" Тексты ниже призваны помочь вам развиваться и предоставляют вам мгновенную оценку вашего прогресса.\n"

	sendMessage := fmt.Sprintf("%s\n📍 %s\n%s\n%s\n%s\n%s", "OZON", location, oby, salary, time, account)

	return sendMessage
}

type service interface {
	CreateLadyAccount(ctx context.Context, account *models.Account) error
	GetAccountLady(ctx context.Context, userName string) (*models.Account, error)
	CreateDonAccount(ctx context.Context, username string) error
	GetDonAccount(ctx context.Context, userName string) (*models.DonAccount, error)
	CreateReview(ctx context.Context, ladyName string, donName string, review string, rating int) error
	GetLadyReviews(ctx context.Context, ladyName string) ([]models.Review, error)
}
