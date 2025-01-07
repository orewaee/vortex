package middlewares

import (
	"github.com/orewaee/vortex/internal/app/domain"
	"github.com/orewaee/vortex/internal/utils"
	"net/http"
)

func PermMiddleware(next http.Handler, permGroup *domain.PermGroup) http.Handler {
	middleware := func(writer http.ResponseWriter, request *http.Request) {
		perms, ok := request.Context().Value("perms").(int)
		if !ok {
			utils.MustWriteString(writer, "missing permissions", http.StatusUnauthorized)
			return
		}

		switch permGroup.GroupMode {
		case domain.GroupModeAll:
			ok = true

			for _, perm := range permGroup.Perms {
				if !domain.HasPerm(perms, perm) {
					ok = false
				}
			}

			if !ok {
				utils.MustWriteString(writer, "permission denied", http.StatusForbidden)
				return
			}
			break

		case domain.GroupModeAny:
			ok = false

			for _, perm := range permGroup.Perms {
				if domain.HasPerm(perms, perm) {
					ok = true
				}
			}

			if !ok {
				utils.MustWriteString(writer, "permission denied", http.StatusForbidden)
				return
			}
			break
		}

		next.ServeHTTP(writer, request)
	}

	return http.HandlerFunc(middleware)
}
