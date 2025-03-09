package rest

import (
	"context"
	"github.com/olahol/melody"
	"github.com/orewaee/vortex/internal/app/api"
	"github.com/orewaee/vortex/internal/app/driving"
	"github.com/orewaee/vortex/internal/cors"
	"github.com/orewaee/vortex/internal/handlers"
	"github.com/rs/zerolog"
	"net/http"
	"sync/atomic"
)

type Controller struct {
	addr   string
	server *http.Server
	melody *melody.Melody

	counter        atomic.Int64
	ticketSessions map[string]map[int64]*melody.Session

	authApi   api.AuthApi
	tokenApi  api.TokenApi
	ticketApi api.TicketApi
	chatApi   api.ChatApi
	log       *zerolog.Logger
}

func NewController(
	addr string,
	authApi api.AuthApi,
	tokenApi api.TokenApi,
	ticketApi api.TicketApi,
	chatApi api.ChatApi,
	log *zerolog.Logger) driving.Controller {
	controller := &Controller{
		addr:           addr,
		melody:         melody.New(),
		ticketSessions: make(map[string]map[int64]*melody.Session),
		authApi:        authApi,
		tokenApi:       tokenApi,
		ticketApi:      ticketApi,
		chatApi:        chatApi,
		log:            log,
	}

	controller.melody.Config.MaxMessageSize = 1024 * 1024 * 8

	return controller
}

func (controller *Controller) Run() error {
	mux := http.NewServeMux()

	mux.HandleFunc("OPTIONS /*", func(writer http.ResponseWriter, _ *http.Request) {
		writer.WriteHeader(http.StatusOK)
	})

	mux.Handle("GET /ping", &handlers.PingHandler{})

	mux.Handle("/v1/", controller.MuxV1())

	controller.server = &http.Server{
		Addr:    controller.addr,
		Handler: cors.NewDefault().Middleware(mux),
	}

	controller.log.Info().Msgf("running on %s", controller.addr)
	return controller.server.ListenAndServe()
}

func (controller *Controller) Shutdown(ctx context.Context) error {
	return controller.server.Shutdown(ctx)
}
