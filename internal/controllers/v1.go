package controllers

import (
	"fmt"
	"github.com/olahol/melody"
	"github.com/orewaee/vortex/internal/app/domain"
	"github.com/orewaee/vortex/internal/handlers"
	"github.com/orewaee/vortex/internal/middlewares"
	"log"
	"net/http"
	"strconv"
	"sync/atomic"
)

func (controller *RestController) MuxV1() http.Handler {
	v1 := http.NewServeMux()

	v1.Handle("POST /login", handlers.NewLoginHandler(controller.authApi))
	v1.HandleFunc("OPTIONS /login", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
	})
	v1.Handle("POST /refresh", handlers.NewRefreshHandler(controller.authApi))
	// v1.Handle("POST /register", nil)

	superGroup := &domain.PermGroup{
		Perms:     []int{domain.PermSuper},
		GroupMode: domain.GroupModeAll,
	}

	v1.Handle("GET /super", middlewares.AuthMiddleware(
		controller.tokenApi,
		middlewares.PermMiddleware(&handlers.SuperHandler{}, superGroup),
	))

	v1.Handle("GET /tickets", middlewares.AuthMiddleware(
		controller.tokenApi,
		middlewares.PermMiddleware(
			handlers.NewTicketsHandler(controller.ticketApi), superGroup),
	))

	v1.Handle("GET /history/{ticket_id}", middlewares.AuthMiddleware(
		controller.tokenApi,
		middlewares.PermMiddleware(
			handlers.NewHistoryHandler(controller.chatApi), superGroup),
	))

	m := melody.New()

	v1.Handle("GET /chat/{ticket_id}", middlewares.AuthMiddleware(
		controller.tokenApi,
		middlewares.PermMiddleware(
			http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
				fmt.Println(request.Method, request.RequestURI)
				m.HandleRequest(writer, request)
			}), superGroup),
	))

	currentSessions := make(map[string]map[int64]*melody.Session)

	go func() {
		messages := controller.chatApi.Subscribe()
		defer controller.chatApi.Unsubscribe(messages)
		controller.log.Info().Msg("listening chat messages in controller...")
		for {
			message := <-messages
			if message.FromSupport {
				continue
			}

			sessions, ok := currentSessions[message.TicketId]
			if !ok {
				continue
			}

			for _, session := range sessions {
				session.Write([]byte(message.Text))
			}
		}
	}()

	var counter atomic.Int64

	m.HandleConnect(func(session *melody.Session) {
		id := counter.Add(1)
		session.Set("id", id)

		ticketId := session.Request.PathValue("ticket_id")

		_, err := controller.ticketApi.GetTicketById(session.Request.Context(), ticketId, false)
		if err != nil {
			log.Println("HANDLE CONNECT", err)
			session.Close()
			return
		}

		_, ok := currentSessions[ticketId]
		if !ok {
			currentSessions[ticketId] = make(map[int64]*melody.Session)
		}

		currentSessions[ticketId][id] = session
	})

	m.HandleMessage(func(session *melody.Session, message []byte) {
		ctx := session.Request.Context()

		name := fmt.Sprintf("%s", ctx.Value("name"))
		ticketId := session.Request.PathValue("ticket_id")

		controller.chatApi.SendMessage(ctx, name, true, ticketId, string(message))
	})

	m.HandleDisconnect(func(session *melody.Session) {
		id, err := strconv.Atoi(fmt.Sprintf("%d", session.MustGet("id")))
		if err != nil {
			panic(err)
		}
		ticketId := session.Request.PathValue("ticket_id")

		delete(currentSessions[ticketId], int64(id))
	})

	return http.StripPrefix("/v1", v1)
}
