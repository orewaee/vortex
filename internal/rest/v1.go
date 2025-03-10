package rest

import (
	"net/http"
)

func (controller *Controller) MuxV1() http.Handler {
	v1 := http.NewServeMux()

	/*
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
	*/

	v1.HandleFunc("GET /tickets", controller.getTickets)

	/*
		v1.Handle("GET /history/{ticket_id}", middlewares.AuthMiddleware(
			controller.tokenApi,
			middlewares.PermMiddleware(
				handlers.NewHistoryHandler(controller.chatApi), superGroup),
		))
	*/

	v1.HandleFunc("GET /chat/{ticket_id}", controller.getChat)
	v1.HandleFunc("GET /chat/history/{ticket_id}", controller.getChatHistory)

	go func() {
		messages := controller.chatApi.Subscribe()
		defer controller.chatApi.Unsubscribe(messages)
		controller.log.Info().Msg("listening chat messages in controller...")
		for {
			message := <-messages
			if message.FromSupport {
				continue
			}

			sessions, ok := controller.ticketSessions[message.TicketId]
			if !ok {
				continue
			}

			for _, session := range sessions {
				session.Write([]byte(message.Text))
			}
		}
	}()

	controller.melody.HandleConnect(controller.handleConnect)
	controller.melody.HandleMessage(controller.handleMessage)
	controller.melody.HandleDisconnect(controller.handleDisconnect)

	return http.StripPrefix("/v1", v1)
}
