package models

type Message struct {
	Key        string `json:"key"`
	Ciphertext string `json:"ciphertext"`
}
