package middlewares

import (
	"context"
	"fmt"
	"github.com/orewaee/typedenv"
	"github.com/orewaee/vortex/internal/app/api"
	"github.com/orewaee/vortex/internal/utils"
	"net/http"
	"strconv"
	"strings"
)

func AuthMiddleware(tokenService api.TokenApi, next http.Handler) http.Handler {
	middleware := func(writer http.ResponseWriter, request *http.Request) {
		header := request.Header.Get("Authorization")
		token := strings.TrimPrefix(header, "Bearer ")

		if token == "" {
			utils.MustWriteString(writer, "missing token", http.StatusUnauthorized)
			return
		}

		claims, err := tokenService.GetTokenClaims(token, typedenv.String("ACCESS_KEY"))
		if err != nil {
			utils.MustWriteString(writer, "invalid token", http.StatusUnauthorized)
			return
		}

		permsClaim, ok := claims["perms"]
		if !ok {
			utils.MustWriteString(writer, "invalid token", http.StatusUnauthorized)
			return
		}

		perms, err := strconv.Atoi(fmt.Sprintf("%v", permsClaim))
		if err != nil {
			utils.MustWriteString(writer, "invalid token", http.StatusUnauthorized)
			return
		}

		nameClaim, ok := claims["name"]
		if !ok {
			utils.MustWriteString(writer, "invalid token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(context.WithValue(request.Context(), "name", nameClaim), "perms", perms)

		next.ServeHTTP(writer, request.WithContext(ctx))
	}

	return http.HandlerFunc(middleware)
}
