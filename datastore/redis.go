package datastore

import (
	"strconv"
	"time"

	"github.com/garyburd/redigo/redis"
)

func redisNewPool(port int) *redis.Pool {
	srv := ":" + strconv.Itoa(port)
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", srv)
			if err != nil {
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}
