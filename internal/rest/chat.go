package rest

import (
	"net/http"
)

func (controller *Controller) getChat(writer http.ResponseWriter, request *http.Request) {
	controller.melody.HandleRequest(writer, request)
}
