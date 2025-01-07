package redis

import (
	"context"
	goredis "github.com/redis/go-redis/v9"
)

func NewClient(ctx context.Context, addr, password string, db int) (*goredis.Client, error) {
	options := &goredis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	}

	client := goredis.NewClient(options)

	if _, err := client.Ping(ctx).Result(); err != nil {
		return nil, err
	}

	return client, nil
}
