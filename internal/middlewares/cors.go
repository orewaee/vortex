package middlewares

import (
	"net/http"
)

func CorsMiddleware(next http.Handler) http.Handler {
	middleware := func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, User-agent")

		next.ServeHTTP(writer, request)
	}

	return http.HandlerFunc(middleware)
}
