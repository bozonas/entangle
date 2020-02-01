package models

type Message struct {
	Key  string `redis:"key"`
	Body string `redis:"body"`
}
