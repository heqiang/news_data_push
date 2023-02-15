package config

import "github.com/go-redis/redis"

var (
	Config  *ServerConfig
	RedisDb *redis.Client
)
