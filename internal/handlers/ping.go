package handlers

import (
	"github.com/orewaee/vortex/internal/utils"
	"net/http"
)

type PingHandler struct{}

func (handler *PingHandler) ServeHTTP(writer http.ResponseWriter, _ *http.Request) {
	utils.MustWriteString(writer, "pong", http.StatusOK)
}
