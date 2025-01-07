package handlers

import (
	"github.com/orewaee/vortex/internal/utils"
	"net/http"
)

type SuperHandler struct{}

func NewSuperHandler() *SuperHandler {
	return &SuperHandler{}
}

func (handler *SuperHandler) ServeHTTP(writer http.ResponseWriter, _ *http.Request) {
	utils.MustWriteString(writer, "super duper secret message", http.StatusOK)
}
