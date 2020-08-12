package db

import (
	"log"
	"time"

	"github.com/garyburd/redigo/redis"
)

var redpool *redis.Pool

const redisurl = "127.0.0.1:6379"

func init() {
	redpool = &redis.Pool{
		MaxIdle:     100,
		MaxActive:   4000,
		Wait:        true,
		IdleTimeout: 180 * time.Second,
		Dial: func() (conn redis.Conn, err error) {
			conn1, err1 := redis.Dial("tcp", redisurl)
			if err != nil {
				log.Println("redis init error ", err)
			}
			return conn1, err1
		},
	}
}

func NewRedis() redis.Conn {
	return redpool.Get()
}
