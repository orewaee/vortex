package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/orewaee/typedenv"
	"github.com/orewaee/vortex/internal/app/api"
	"github.com/orewaee/vortex/internal/config"
	"github.com/orewaee/vortex/internal/logger"
	"github.com/orewaee/vortex/internal/repo/postgres"
	"github.com/orewaee/vortex/internal/repo/redis"
	"github.com/orewaee/vortex/internal/rest"
	"github.com/orewaee/vortex/internal/services"
	"github.com/orewaee/vortex/internal/telegram"
	goredis "github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
)

func main() {
	config.MustLoad()

	ctx := context.Background()

	postgresPool := mustInitPostgresPool(ctx)
	redisClient := mustInitRedisClient(ctx)

	log, err := logger.NewZerolog()
	if err != nil {
		panic(err)
	}

	tokenApi := mustInitTokenApi(redisClient)
	authApi := services.NewAuthService(tokenApi, log)
	ticketApi := mustInitTicketApi(postgresPool, log)
	chatApi := mustInitChatApi(postgresPool, log)

	token := typedenv.String("TELEGRAM_TOKEN")
	telegramBot := telegram.NewBot(token, ticketApi, chatApi, log)
	go telegramBot.Run()

	addr := typedenv.String("VORTEX_ADDR")
	controller := rest.NewController(addr, authApi, tokenApi, ticketApi, chatApi, log)

	if err := controller.Run(); err != nil {
		panic(err)
	}
}

func mustInitTokenApi(client *goredis.Client) api.TokenApi {
	tokenRepo := redis.NewTokenRepo(client)
	return services.NewJwtTokenService(tokenRepo)
}

func mustInitTicketApi(pool *pgxpool.Pool, log *zerolog.Logger) api.TicketApi {
	repo := postgres.NewTicketRepo(pool)
	return services.NewTicketService(repo, log)
}

func mustInitChatApi(pool *pgxpool.Pool, log *zerolog.Logger) api.ChatApi {
	repo := postgres.NewChatRepo(pool)
	return services.NewChatService(repo, log)
}

func mustInitPostgresPool(ctx context.Context) *pgxpool.Pool {
	user := typedenv.String("POSTGRES_USER")
	password := typedenv.String("POSTGRES_PASSWORD")
	host := typedenv.String("POSTGRES_HOST")
	port := typedenv.String("POSTGRES_PORT")

	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/vortex?sslmode=disable", user, password, host, port)

	pool, err := postgres.NewPool(ctx, connString)
	if err != nil {
		panic(err)
	}

	return pool
}

func mustInitRedisClient(ctx context.Context) *goredis.Client {
	host := typedenv.String("REDIS_HOST")
	port := typedenv.String("REDIS_PORT")
	password := typedenv.String("REDIS_PASSWORD")

	addr := fmt.Sprintf("%s:%s", host, port)

	client, err := redis.NewClient(ctx, addr, password, 0)
	if err != nil {
		panic(err)
	}

	return client
}
