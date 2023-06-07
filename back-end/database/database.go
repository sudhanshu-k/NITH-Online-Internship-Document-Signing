package database

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

// DB pgx connector
var DB *pgxpool.Pool

// Reddis client
var RedisClient *redis.Client
