package initializers

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing/tree/main/back-end/utils"
)

var (
	RedisClient *redis.Client
	ctx         context.Context
)

func ConnectRedis(config *Config) {
	ctx = context.TODO()

	RedisClient = redis.NewClient(&redis.Options{
		Addr:      "redis-10241.c281.us-east-1-2.ec2.cloud.redislabs.com:10241",
		Username:  "default",
		Password:  "1J6kFiBoAzGXvYZM5qIb2VBBoTzmagwn",
		// TLSConfig: &tls.Config{MinVersion: tls.VersionTLS12}, // define your TLS configuration here
	})

	_, err := RedisClient.Ping(ctx).Result()
	utils.FatalError(err)

	err = RedisClient.Set(ctx, "API", "e-sign application", 0).Err()
	utils.FatalError(err)

	fmt.Println("Redis Connected")
}
