package rest

import (
	"context"
	"github.com/orewaee/vortex/internal/app/api"
	"github.com/orewaee/vortex/internal/app/driving"
	"github.com/orewaee/vortex/internal/cors"
	"github.com/orewaee/vortex/internal/handlers"
	"github.com/rs/zerolog"
	"net/http"
)

type Controller struct {
	addr      string
	server    *http.Server
	authApi   api.AuthApi
	tokenApi  api.TokenApi
	ticketApi api.TicketApi
	chatApi   api.ChatApi
	log       *zerolog.Logger
}

func NewController(
	addr string,
	server *http.Server,
	authApi api.AuthApi,
	tokenApi api.TokenApi,
	ticketApi api.TicketApi,
	chatApi api.ChatApi,
	log *zerolog.Logger) driving.Controller {
	return &Controller{
		addr:      addr,
		server:    server,
		authApi:   authApi,
		tokenApi:  tokenApi,
		ticketApi: ticketApi,
		chatApi:   chatApi,
		log:       log,
	}
}

func (controller *Controller) Run() error {
	mux := http.NewServeMux()

	optionsHandler := func(writer http.ResponseWriter, _ *http.Request) {
		writer.WriteHeader(http.StatusOK)
	}

	mux.HandleFunc("OPTIONS /*", optionsHandler)

	mux.Handle("GET /ping", &handlers.PingHandler{})

	mux.Handle("/v1/", controller.MuxV1())

	server := &http.Server{
		Addr:    controller.addr,
		Handler: cors.NewDefault().Middleware(mux),
	}

	controller.log.Info().Msgf("running on %s", controller.addr)
	return server.ListenAndServe()
}

func (controller *Controller) Shutdown(ctx context.Context) error {
	return controller.server.Shutdown(ctx)
}
