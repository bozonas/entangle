package main

import (
	"errors"

	"github.com/gomodule/redigo/redis"
)

var pool *redis.Pool

var ErrNoMessage = errors.New("no message found")

type Message struct {
	Key  string `redis:"key"`
	Body string `redis:"body"`
}

func FindMessage(key string) (*Message, error) {
	conn := pool.Get()

	defer conn.Close()

	values, err := redis.Values(conn.Do("HGETALL", "message:"+key))
	if err != nil {
		return nil, err
	} else if len(values) == 0 {
		return nil, ErrNoMessage
	}

	var message Message
	err = redis.ScanStruct(values, &message)
	if err != nil {
		return nil, err
	}

	return &message, nil
}
