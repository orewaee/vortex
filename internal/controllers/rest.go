package controllers

import (
	"github.com/orewaee/vortex/internal/app/api"
	"github.com/orewaee/vortex/internal/handlers"
	"github.com/orewaee/vortex/internal/middlewares"
	"github.com/rs/zerolog"
	"net/http"
)

type RestController struct {
	authApi   api.AuthApi
	tokenApi  api.TokenApi
	ticketApi api.TicketApi
	chatApi   api.ChatApi
	log       *zerolog.Logger
}

func NewRestController(
	authApi api.AuthApi, tokenApi api.TokenApi,
	ticketApi api.TicketApi, chatApi api.ChatApi,
	log *zerolog.Logger) *RestController {
	return &RestController{
		authApi:   authApi,
		tokenApi:  tokenApi,
		ticketApi: ticketApi,
		chatApi:   chatApi,
		log:       log,
	}
}

func (controller *RestController) Run(addr string) error {
	mux := http.NewServeMux()

	mux.Handle("GET /ping", &handlers.PingHandler{})

	mux.Handle("/v1/", middlewares.CorsMiddleware(controller.MuxV1()))

	server := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	controller.log.Info().Msgf("running on %s", addr)
	return server.ListenAndServe()
}
