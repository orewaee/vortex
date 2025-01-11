package handlers

import (
	"github.com/orewaee/vortex/internal/app/api"
	"github.com/orewaee/vortex/internal/dtos"
	"github.com/orewaee/vortex/internal/utils"
	"net/http"
	"strconv"
)

type HistoryHandler struct {
	chatApi api.ChatApi
}

func NewHistoryHandler(chatApi api.ChatApi) http.Handler {
	return &HistoryHandler{chatApi}
}

func (handler *HistoryHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	ticketId := request.PathValue("ticket_id")

	limit := 20
	queryLimit := request.URL.Query().Get("limit")
	if queryLimit != "" {
		temp, err := strconv.Atoi(queryLimit)
		if err == nil {
			limit = temp
		}
	}

	offset := 0
	queryOffset := request.URL.Query().Get("offset")
	if queryOffset != "" {
		temp, err := strconv.Atoi(queryLimit)
		if err == nil {
			offset = temp
		}
	}

	messages, err := handler.chatApi.GetMessageHistory(request.Context(), ticketId, limit, offset)
	if err != nil {
		utils.MustWriteString(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	data := make([]*dtos.Message, len(messages))
	for i, message := range messages {
		data[i] = &dtos.Message{
			Id:          message.Id,
			Sender:      message.Sender,
			FromSupport: message.FromSupport,
			TicketId:    message.TicketId,
			Text:        message.Text,
			CreatedAt:   message.CreatedAt,
		}
	}

	utils.MustWriteJson(writer, data, http.StatusOK)
}
