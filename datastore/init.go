package datastore

import (
	"log"

	"github.com/garyburd/redigo/redis"
)

var redisPool *redis.Pool

func Initialize(redisPort int) {
	redisPool = redisNewPool(redisPort)
}

// must be run in separate goroutine
func Subscribe(channel string, msg chan<- []byte) {
	c := redisPool.Get()
	p := redis.PubSubConn{c}
	p.Subscribe(channel)
	for {
		switch v := p.Receive().(type) {
		case redis.Message:
			msg <- v.Data
		case redis.Subscription:
			log.Printf("%s: %s %d\n", v.Channel, v.Kind, v.Count)
		case error:
			log.Printf("%v\n", v)
		}
	}
}
