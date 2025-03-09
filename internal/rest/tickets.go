package rest

import (
	"github.com/orewaee/vortex/internal/dtos"
	"github.com/orewaee/vortex/internal/utils"
	"net/http"
	"strconv"
)

func (controller *Controller) getTickets(writer http.ResponseWriter, request *http.Request) {
	page := 0
	queryPage := request.URL.Query().Get("page")
	if queryPage != "" {
		temp, err := strconv.Atoi(queryPage)
		if err == nil {
			page = temp
		}
	}

	perPage := 20
	queryPerPage := request.URL.Query().Get("per_page")
	if queryPerPage != "" {
		temp, err := strconv.Atoi(queryPerPage)
		if err == nil {
			perPage = temp
		}
	}

	tickets, err := controller.ticketApi.GetTickets(request.Context(), page, perPage)
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
