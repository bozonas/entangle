package middleware

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"../models"
	"github.com/gomodule/redigo/redis"
	"github.com/gorilla/mux"
)

func init() {
	pool = &redis.Pool{
		MaxIdle:     10,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379")
		},
	}

	// models.Posts = append(models.Posts, Post{ID: "1", Title: "My first post", Body: "This is the content of my first post"})
}

func GetPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.Posts)
}

func GetPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range models.Posts {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			break
		}
		return
	}
	json.NewEncoder(w).Encode(&models.Post{})
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range models.Posts {
		if item.ID == params["id"] {
			models.Posts = append(models.Posts[:index], models.Posts[index+1:]...)
			var post models.Post
			_ = json.NewDecoder(r.Body).Decode(post)
			post.ID = params["id"]
			models.Posts = append(models.Posts, post)
			json.NewEncoder(w).Encode(&post)
			return
		}
	}
	json.NewEncoder(w).Encode(models.Posts)
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range models.Posts {
		if item.ID == params["id"] {
			models.Posts = append(models.Posts[:index], models.Posts[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(models.Posts)
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var post models.Post
	_ = json.NewDecoder(r.Body).Decode(post)
	post.ID = strconv.Itoa(rand.Intn(1000000))
	models.Posts = append(models.Posts, post)
	json.NewEncoder(w).Encode(&post)
}

func SaveMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var message models.Message
	_ = json.NewDecoder(r.Body).Decode(message)
	// post.ID = strconv.Itoa(rand.Intn(1000000))
	// models.Posts = append(models.Posts, post)
	// json.NewEncoder(w).Encode(&message)
}

func GetMessage(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if id == "" {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	message, err := FindMessage(id)
	if err == ErrNoMessage {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(message)
}

func SetMessage(w http.ResponseWriter, r *http.Request) {

}
