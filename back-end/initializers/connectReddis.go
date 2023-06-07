package initializers

import (
	"context"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"

	"github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing/tree/main/back-end/config"
	"github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing/tree/main/back-end/database"
	"github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing/tree/main/back-end/utils"
)

func ConnectRedis() {
	ctx := context.TODO()

	// Local host connection
	database.RedisClient = redis.NewClient(&redis.Options{
		Addr: config.Config.RedisURL,
	})

	_, err := database.RedisClient.Ping(ctx).Result()
	if err != nil {
		utils.Logger.Fatal("Unable to connect to reddis.", zap.Error(err))
	}

	err = database.RedisClient.Set(ctx, "API", "e-sign application", 0).Err()
	if err != nil {
		utils.Logger.Fatal("Error occured while executing test query.", zap.Error(err))
	}

	utils.Logger.Info("Connected to reddis.")
}
