package bot

import (
	"context"
	"errors"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/orewaee/vortex/internal/app/api"
	"github.com/orewaee/vortex/internal/app/domain"
	"github.com/rs/zerolog"
	"strings"
)

type Bot struct {
	botApi    *tgbotapi.BotAPI
	ticketApi api.TicketApi
	chatApi   api.ChatApi
	log       *zerolog.Logger
}

func NewBot(token string, ticketApi api.TicketApi, chatApi api.ChatApi, log *zerolog.Logger) *Bot {
	botApi, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		panic(err)
	}

	return &Bot{
		botApi:    botApi,
		ticketApi: ticketApi,
		chatApi:   chatApi,
		log:       log,
	}
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
		bot.log.Info().Msg("listening chat messages in bot...")

		for {
			message := <-messages
			if message.FromSupport {
				ticket, err := bot.ticketApi.GetTicketById(ctx, message.TicketId, false)
				if err != nil || ticket == nil {
					continue
				}

				msg := tgbotapi.NewMessage(ticket.ChatId, message.Text)
				msg.ParseMode = tgbotapi.ModeMarkdown

				bot.botApi.Send(msg)
			}
		}
	}()

	for update := range updates {
		message := update.Message
		query := update.CallbackQuery

		if query != nil && query.Data == "close_ticket" {
			callback := tgbotapi.NewCallback(query.ID, "Processing...")
			if _, err := bot.botApi.Request(callback); err != nil {
				panic(err)
			}

			chatId := query.Message.Chat.ID

			ticket, err := bot.ticketApi.GetTicketByChatId(ctx, chatId, false)
			if ticket == nil || err != nil {
				continue
			}

			bot.ticketApi.CloseTicketById(ctx, ticket.Id)

			text := fmt.Sprintf("Ticket `%s` has been closed.", ticket.Id)
			edit := tgbotapi.NewEditMessageText(chatId, query.Message.MessageID, text)
			edit.ParseMode = tgbotapi.ModeMarkdown
			bot.botApi.Send(edit)

			continue
		}

		if message == nil {
			continue
		}

		if strings.HasPrefix(message.Text, "/ticket") {
			topic := message.CommandArguments()

			ticket, err := bot.ticketApi.GetTicketByChatId(ctx, message.Chat.ID, false)
			if err != nil && !errors.Is(err, domain.ErrTicketNotFound) {
				bot.log.Err(err).Send()
				continue
			}

			if err == nil {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "You already have an open ticket.")
				msg.ReplyToMessageID = update.Message.MessageID
				msg.ParseMode = tgbotapi.ModeMarkdown

				msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
					tgbotapi.NewInlineKeyboardRow(
						tgbotapi.NewInlineKeyboardButtonData("Close", "close_ticket"),
					),
				)

				bot.botApi.Send(msg)
				continue
			}

			ticket, err = bot.ticketApi.OpenTicket(ctx, message.Chat.ID, topic)

			text := ""
			if err != nil {
				text = "error. " + err.Error()
			} else {
				text = fmt.Sprintf("Ticket's open. Its id: `%s`.", ticket.Id)
			}

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
			msg.ReplyToMessageID = update.Message.MessageID
			msg.ParseMode = tgbotapi.ModeMarkdown

			msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("Close", "close_ticket"),
				),
			)

			bot.botApi.Send(msg)
			continue
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

			bot.botApi.Send(msg)
			continue
		}

		ticket, err := bot.ticketApi.GetTicketByChatId(ctx, message.Chat.ID, false)
		if err != nil || ticket == nil {
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
