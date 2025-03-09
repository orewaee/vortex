package rest

import (
	"fmt"
	"github.com/olahol/melody"
	"log"
	"strconv"
)

func (controller *Controller) handleConnect(session *melody.Session) {
	id := controller.counter.Add(1)
	session.Set("id", id)

	ticketId := session.Request.PathValue("ticket_id")

	ticket, err := controller.ticketApi.GetTicketById(session.Request.Context(), ticketId)
	if err != nil {
		log.Println("HANDLE CONNECT", err)
		session.Close()
		return
	}

	fmt.Println("CONNECTION TO THE", ticket)

	_, ok := controller.ticketSessions[ticketId]
	if !ok {
		controller.ticketSessions[ticketId] = make(map[int64]*melody.Session)
	}

	controller.ticketSessions[ticketId][id] = session
}

func (controller *Controller) handleMessage(session *melody.Session, message []byte) {
	ctx := session.Request.Context()

	name := fmt.Sprintf("%s", ctx.Value("name"))
	ticketId := session.Request.PathValue("ticket_id")

	controller.chatApi.SendMessage(ctx, name, true, ticketId, string(message))
}

func (controller *Controller) handleDisconnect(session *melody.Session) {
	id, err := strconv.Atoi(fmt.Sprintf("%d", session.MustGet("id")))
	if err != nil {
		panic(err)
	}
	ticketId := session.Request.PathValue("ticket_id")

	delete(controller.ticketSessions[ticketId], int64(id))
}
