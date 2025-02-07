package cors

import (
	"net/http"
	"strings"
)

type Cors struct {
	config *Config
}

func New(config *Config) *Cors {
	return &Cors{config: config}
}

func NewDefault() *Cors {
	return New(DefaultConfig())
}

func (cors *Cors) Middleware(next http.Handler) http.Handler {
	middleware := func(writer http.ResponseWriter, request *http.Request) {
		setHeader := func(key string, values []string) {
			writer.Header().Set(key, strings.Join(values, ","))
		}

		setHeader("Access-Control-Allow-Origin", cors.config.AllowedOrigins)
		setHeader("Access-Control-Allow-Methods", cors.config.AllowedMethods)
		setHeader("Access-Control-Allow-Headers", cors.config.AllowedHeaders)

		next.ServeHTTP(writer, request)
	}

	return http.HandlerFunc(middleware)
}
