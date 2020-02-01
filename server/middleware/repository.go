package middleware

import (
	"errors"

	"../models"

	"github.com/gomodule/redigo/redis"
)

var pool *redis.Pool

var ErrNoMessage = errors.New("no message found")

func FindMessage(key string) (*models.Message, error) {
	conn := pool.Get()

	defer conn.Close()

	values, err := redis.Values(conn.Do("HGETALL", "message:"+key))
	if err != nil {
		return nil, err
	} else if len(values) == 0 {
		return nil, ErrNoMessage
	}

	var message models.Message
	err = redis.ScanStruct(values, &message)
	if err != nil {
		return nil, err
	}

	return &message, nil
}
