package handlers

import (
	"github.com/orewaee/vortex/internal/app/api"
	"github.com/orewaee/vortex/internal/dtos"
	"github.com/orewaee/vortex/internal/utils"
	"net/http"
	"strconv"
)

type TicketsHandler struct {
	ticketApi api.TicketApi
}

func NewTicketsHandler(ticketApi api.TicketApi) *TicketsHandler {
	return &TicketsHandler{ticketApi}
}

func (handler *TicketsHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
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

	tickets, err := handler.ticketApi.GetTickets(request.Context(), limit, offset)
	if err != nil {
		utils.MustWriteString(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	data := make([]*dtos.Ticket, len(tickets))
	for i, ticket := range tickets {
		data[i] = &dtos.Ticket{
			Id:        ticket.Id,
			ChatId:    ticket.ChatId,
			Topic:     ticket.Topic,
			Closed:    ticket.Closed,
			CreatedAt: ticket.CreatedAt,
		}
	}

	utils.MustWriteJson(writer, data, http.StatusOK)
}
