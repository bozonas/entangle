package middleware

import (
	"encoding/json"
	"errors"
	"flag"
	"log"
	"os"
	"time"

	"entangle/models"

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
	var configuration Configuration
	if *filename == "-" {
		configuration = Configuration{
			RedisUrl:      os.Getenv("PORT"),
			RedisPassword: os.Getenv("PORT"),
		}
	} else {
		file, _ := os.Open(*filename)
		decoder := json.NewDecoder(file)
		decoder.Decode(&configuration)
	}

	pool = &redis.Pool{
		MaxIdle:     10,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp",
				"redis-13701.c98.us-east-1-4.ec2.cloud.redislabs.com:13701",
				redis.DialPassword("dq4UvSdmacMaQ0iEVkAaWb3gEPNw7qJI"))
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
