package store

import (
	"os"

	"github.com/gomodule/redigo/redis"
)

func RedisInit() *redis.Pool {
	return &redis.Pool{
		MaxIdle: 10,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", os.Getenv("REDIS"))
		},
	}
}
