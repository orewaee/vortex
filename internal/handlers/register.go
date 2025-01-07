package handlers

import (
	"github.com/orewaee/vortex/internal/utils"
	"net/http"
)

type RegisterHandler struct{}

func (handler *RegisterHandler) ServeHTTP(writer http.ResponseWriter, _ *http.Request) {
	utils.MustWriteString(writer, "unimplemented", http.StatusOK)
}
