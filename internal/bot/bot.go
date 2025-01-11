package bot

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/orewaee/vortex/internal/app/api"
	"log"
	"strings"
)

type Bot struct {
	botApi    *tgbotapi.BotAPI
	ticketApi api.TicketApi
	chatApi   api.ChatApi
}

func NewBot(token string, ticketApi api.TicketApi, chatApi api.ChatApi) *Bot {
	botApi, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		panic(err)
	}

	return &Bot{botApi, ticketApi, chatApi}
}

func (bot *Bot) MustRun() {
	ctx := context.Background()

	bot.botApi.Debug = false

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30
	updates := bot.botApi.GetUpdatesChan(updateConfig)

	go func() {
		messages := bot.chatApi.Subscribe()
		defer bot.chatApi.Unsubscribe(messages)
		log.Println("listening chat messages in bot...")

		for {
			message := <-messages
			if message.FromSupport {
				ticket, err := bot.ticketApi.GetTicketById(ctx, message.TicketId)
				if err != nil {
					continue
				}

				msg := tgbotapi.NewMessage(ticket.ChatId, message.Text)
				msg.ParseMode = tgbotapi.ModeMarkdown

				if _, err := bot.botApi.Send(msg); err != nil {
					panic(err)
				}
			}
		}
	}()

	for update := range updates {
		message := update.Message

		if message == nil {
			continue
		}

		if strings.HasPrefix(message.Text, "/ticket") {
			topic := message.CommandArguments()
			log.Println("topic:", topic)

			ticket, err := bot.ticketApi.OpenTicket(ctx, message.Chat.ID, topic)

			text := ""
			if err != nil {
				text = "error. " + err.Error()
			} else {
				text = fmt.Sprintf("ticket created. your ticket id: `%s`", ticket.Id)
			}

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
			msg.ReplyToMessageID = update.Message.MessageID
			msg.ParseMode = tgbotapi.ModeMarkdown

			if _, err := bot.botApi.Send(msg); err != nil {
				panic(err)
			}
		}

		if message.Text == "/close" {
			err := bot.ticketApi.CloseTicketByChatId(ctx, message.Chat.ID)

			text := ""
			if err != nil {
				text = "error. " + err.Error()
			} else {
				text = fmt.Sprintf("ticket closed.")
			}

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
			msg.ReplyToMessageID = update.Message.MessageID
			msg.ParseMode = tgbotapi.ModeMarkdown

			if _, err := bot.botApi.Send(msg); err != nil {
				panic(err)
			}
		}

		ticket, err := bot.ticketApi.GetTicketByChatId(ctx, message.Chat.ID)
		if err != nil {
			fmt.Println(err)
			continue
		}

		from := message.From
		name := from.UserName
		if name == "" {
			name = fmt.Sprintf("%s %s", from.FirstName, from.LastName)
		}

		bot.chatApi.SendMessage(ctx, name, false, ticket.Id, message.Text)
	}
}
