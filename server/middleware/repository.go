package middleware

import (
	"errors"
	"log"
	"time"

	"../models"

	"github.com/gomodule/redigo/redis"
)

var pool *redis.Pool

var ErrNoMessage = errors.New("no message found")

func init() {
	pool = &redis.Pool{
		MaxIdle:     10,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp",
				"XXXXX",
				redis.DialPassword("XXXXXXX"))
		},
	}
}

func SetMessage(message *models.Message) error {
	conn := pool.Get()
	defer conn.Close()

	// set expiry for 1 hour
	if _, err := conn.Do("SETEX", message.Key, 3600, message.Ciphertext); err != nil {
		log.Fatal(err)
	}

	return nil
}

func FindMessage(key string) (*models.Message, error) {
	conn := pool.Get()
	defer conn.Close()

	str, err := redis.String(conn.Do("GET", key))
	if err != nil {
		log.Fatal(err)
	}

	if _, err := conn.Do("DEL", key); err != nil {
		log.Fatal(err)
	}

	message := &models.Message{
		Key:        key,
		Ciphertext: str,
	}

	return message, nil
}
