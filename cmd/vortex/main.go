package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/orewaee/typedenv"
	"github.com/orewaee/vortex/internal/app/api"
	"github.com/orewaee/vortex/internal/bot"
	"github.com/orewaee/vortex/internal/config"
	"github.com/orewaee/vortex/internal/controllers"
	"github.com/orewaee/vortex/internal/logger"
	"github.com/orewaee/vortex/internal/repo/postgres"
	"github.com/orewaee/vortex/internal/repo/redis"
	"github.com/orewaee/vortex/internal/services"
	goredis "github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
)

func main() {
	config.MustLoad()

	ctx := context.Background()

	postgresPool := mustInitPostgres(ctx)
	redisClient := mustInitRedis(ctx)

	log, err := logger.NewZerolog()
	if err != nil {
		panic(err)
	}

	tokenApi := mustInitTokenApi(redisClient)
	authApi := services.NewAuthService(tokenApi, log)
	ticketApi := mustInitTicketApi(postgresPool, log)
	chatApi := mustInitChatApi(postgresPool, log)

	token := typedenv.String("TELEGRAM_TOKEN")
	telegramBot := bot.NewBot(token, ticketApi, chatApi, log)
	go telegramBot.MustRun()

	rest := controllers.NewRestController(authApi, tokenApi, ticketApi, chatApi, log)

	addr := typedenv.String("VORTEX_ADDR")

	if err := rest.Run(addr); err != nil {
		panic(err)
	}
}

func mustInitTokenApi(client *goredis.Client) api.TokenApi {
	tokenRepo := redis.NewTokenRepo(client)
	return services.NewJwtTokenService(tokenRepo)
}

func mustInitTicketApi(pool *pgxpool.Pool, log *zerolog.Logger) api.TicketApi {
	repo, err := postgres.NewTicketRepo(pool)
	if err != nil {
		panic(err)
	}

	return services.NewTicketService(repo, log)
}

func mustInitChatApi(pool *pgxpool.Pool, log *zerolog.Logger) api.ChatApi {
	repo, err := postgres.NewChatRepo(pool)
	if err != nil {
		panic(err)
	}

	return services.NewChatService(repo, log)
}

func mustInitPostgres(ctx context.Context) *pgxpool.Pool {
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

	return pool
}

func mustInitRedis(ctx context.Context) *goredis.Client {
	addr := typedenv.String("REDIS_ADDR", ":6379")
	password := typedenv.String("REDIS_PASSWORD")

	client, err := redis.NewClient(ctx, addr, password, 0)
	if err != nil {
		panic(err)
	}

	return client
}
