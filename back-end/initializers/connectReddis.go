package initializers

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing/tree/main/back-end/middleware"
)

var (
	RedisClient *redis.Client
	ctx         context.Context
)

func ConnectRedis(config *Config) {
	ctx = context.TODO()

	RedisClient = redis.NewClient(&redis.Options{
		Addr: config.RedisURL,
	})

	_, err := RedisClient.Ping(ctx).Result()
	middleware.LogIfError(err, "Reddis ping failed.")

	err = RedisClient.Set(ctx, "API", "e-sign application", 0).Err()
	middleware.FatalError(err)

	fmt.Println("Redis Connected")
}
