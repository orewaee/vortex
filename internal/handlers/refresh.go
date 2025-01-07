package handlers

import (
	"github.com/orewaee/vortex/internal/app/api"
	"github.com/orewaee/vortex/internal/dtos"
	"github.com/orewaee/vortex/internal/utils"
	"net/http"
)

type RefreshHandler struct {
	authService api.AuthApi
}

func NewRefreshHandler(authService api.AuthApi) *RefreshHandler {
	return &RefreshHandler{authService}
}

func (handler *RefreshHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	data := utils.MustReadJson[dtos.RefreshRequest](request)

	// todo: validate data

	access, refresh, err := handler.authService.Refresh(request.Context(), data.RefreshToken)

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
