package main

import (
	"context"
	"fmt"
	"github.com/orewaee/typedenv"
	"github.com/orewaee/vortex/internal/app/api"
	"github.com/orewaee/vortex/internal/bot"
	"github.com/orewaee/vortex/internal/config"
	"github.com/orewaee/vortex/internal/controllers"
	"github.com/orewaee/vortex/internal/repo/postgres"
	"github.com/orewaee/vortex/internal/repo/redis"
	"github.com/orewaee/vortex/internal/services"
)

func main() {
	config.MustLoad()

	ctx := context.Background()

	tokenService := mustInitTokenService(ctx)
	authService := services.NewAuthService(tokenService)
	ticketService := mustInitTicketService(ctx)

	chatApi := services.NewChatService()

	telegramBot := bot.NewBot(typedenv.String("TELEGRAM_TOKEN"), ticketService, chatApi)
	go telegramBot.MustRun()

	rest := controllers.NewRestController(authService, tokenService, ticketService, chatApi)

	addr := typedenv.String("VORTEX_ADDR")

	if err := rest.Run(addr); err != nil {
		panic(err)
	}
}

func mustInitTokenService(ctx context.Context) api.TokenApi {
	addr := typedenv.String("REDIS_ADDR", ":6379")
	password := typedenv.String("REDIS_PASSWORD")

	client, err := redis.NewClient(ctx, addr, password, 0)
	if err != nil {
		panic(err)
	}

	tokenRepo := redis.NewTokenRepo(client)

	return services.NewJwtTokenService(tokenRepo)
}

func mustInitTicketService(ctx context.Context) api.TicketApi {
	user := typedenv.String("POSTGRES_USER")
	password := typedenv.String("POSTGRES_PASSWORD")
	addr := typedenv.String("POSTGRES_ADDR", ":5432")
	database := typedenv.String("POSTGRES_DATABASE")

	connString := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
		user, password, addr, database)

	pool, err := postgres.NewPool(ctx, connString)
	if err != nil {
		panic(err)
	}

	repo, err := postgres.NewTicketRepo(pool)
	if err != nil {
		panic(err)
	}

	return services.NewTicketService(repo)
}
