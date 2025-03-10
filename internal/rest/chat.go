package rest

import (
	"github.com/orewaee/vortex/internal/dtos"
	"github.com/orewaee/vortex/internal/utils"
	"net/http"
)

func (controller *Controller) getChat(writer http.ResponseWriter, request *http.Request) {
	controller.melody.HandleRequest(writer, request)
}

func (controller *Controller) getChatHistory(writer http.ResponseWriter, request *http.Request) {
	ticketId := request.PathValue("ticket_id")
	page := utils.IntQueryParam(request, "page")
	perPage := utils.IntQueryParam(request, "per_page")

	messages, err := controller.chatApi.GetMessageHistory(request.Context(), ticketId, page, perPage)
	if err != nil {
		controller.log.Error().Err(err).Send()
		utils.MustWriteJson(writer, &dtos.Error{Message: err.Error()}, http.StatusInternalServerError)
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
