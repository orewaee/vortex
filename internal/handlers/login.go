package handlers

import (
	"github.com/orewaee/vortex/internal/app/api"
	"github.com/orewaee/vortex/internal/dtos"
	"github.com/orewaee/vortex/internal/utils"
	"net/http"
)

type LoginHandler struct {
	authService api.AuthApi
}

func NewLoginHandler(authService api.AuthApi) *LoginHandler {
	return &LoginHandler{authService}
}

func (handler *LoginHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	data := utils.MustReadJson[dtos.LoginRequest](request)

	// todo: validate data

	access, refresh, err := handler.authService.Login(request.Context(), data.Name, data.Password)

	if err != nil {
		utils.MustWriteString(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	pair := &dtos.TokenPair{
		AccessToken:  access,
		RefreshToken: refresh,
	}

	utils.MustWriteJson(writer, pair, http.StatusOK)
}
