package middleware

import (
	"encoding/json"
	"errors"
	"flag"
	"log"
	"os"
	"time"

	"../models"

	"github.com/gomodule/redigo/redis"
)

var pool *redis.Pool

var ErrNoMessage = errors.New("no message found")

var filename = flag.String("config", "config.json", "Location of the config file.")

type Configuration struct {
	RedisUrl      string
	RedisPassword string
}

func init() {
	flag.Parse()
	file, _ := os.Open(*filename)
	var configuration Configuration
	decoder := json.NewDecoder(file)
	decoder.Decode(&configuration)

	pool = &redis.Pool{
		MaxIdle:     10,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp",
				configuration.RedisUrl,
				redis.DialPassword(configuration.RedisPassword))
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
		log.Print(err)
		return nil, ErrNoMessage
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
