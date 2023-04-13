package initializers

import (
	"context"

	"github.com/redis/go-redis/v9"
	"github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing/tree/main/back-end/utils"

)

var (
	RedisClient *redis.Client
	ctx         context.Context
)

func ConnectRedis(config *Config) (*redis.Client, error){
	// ctx = context.TODO()

	// rdb := redis.NewClient(&redis.Options{
	// 	Addr: "mSg3XP306sYfwOc3yo5FMgKfQ3z1Fvg6@oregon-redis.render.com:6379", // Replace with your Redis server address
	// 	// rediss://red-cgs3hiss3fvk98282sbg:mSg3XP306sYfwOc3yo5FMgKfQ3z1Fvg6@oregon-redis.render.com:6379
	// 	Password: "", // Replace with your Redis server password, if applicable
	// 	DB:       0,
	// })

	// // RedisClient = redis.NewClient(&redis.Options{
	// // 	Addr: config.RedisURL,
	// // })
	// // fmt.Println("sdf")
	// // _, err := rdb.Ping(ctx).Result()
	// // utils.FatalError(err)

	// err := rdb.Set(ctx, "API", "e-sign application", 0).Err()
	// utils.FatalError(err)

	// fmt.Println("Redis Connected")4
	ctx = context.Background()
	redisURL := "rediss://red-cgs3hiss3fvk98282sbg:mSg3XP306sYfwOc3yo5FMgKfQ3z1Fvg6@oregon-redis.render.com:6379"

	options, err := redis.ParseURL(redisURL)
	if err != nil {
		return nil, err
	}

	client := redis.NewClient(options)

	// Ping the Redis server to verify that the connection was successful
	if err := client.Ping(ctx).Err(); err != nil {
		return nil, err
	}

		err = client.Set(ctx, "API", "e-sign application", 0).Err()
	utils.FatalError(err)

	return client , nil
}
