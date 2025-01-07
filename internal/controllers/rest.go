package controllers

import (
	"github.com/orewaee/vortex/internal/app/api"
	"github.com/orewaee/vortex/internal/handlers"
	"log"
	"net/http"
)

type RestController struct {
	authService  api.AuthApi
	tokenService api.TokenApi
	ticketApi    api.TicketApi
	chatApi      api.ChatApi
}

func NewRestController(authService api.AuthApi, tokenService api.TokenApi, ticketApi api.TicketApi, chatApi api.ChatApi) *RestController {
	return &RestController{authService, tokenService, ticketApi, chatApi}
}

func (controller *RestController) Run(addr string) error {
	mux := http.NewServeMux()

	mux.Handle("GET /ping", &handlers.PingHandler{})

	mux.Handle("/v1/", controller.MuxV1())

	server := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	log.Printf("running on %s\n", addr)

	return server.ListenAndServe()
}
