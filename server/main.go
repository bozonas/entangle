package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/gorilla/mux"
)

type Post struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

// type Message struct {
// 	Id   int    `json:"id"`
// 	Body string `json:"body"`
// }

var posts []Post

func getPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

func getPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range posts {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			break
		}
		return
	}
	json.NewEncoder(w).Encode(&Post{})
}

func updatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range posts {
		if item.ID == params["id"] {
			posts = append(posts[:index], posts[index+1:]...)
			var post Post
			_ = json.NewDecoder(r.Body).Decode(post)
			post.ID = params["id"]
			posts = append(posts, post)
			json.NewEncoder(w).Encode(&post)
			return
		}
	}
	json.NewEncoder(w).Encode(posts)
}

func deletePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range posts {
		if item.ID == params["id"] {
			posts = append(posts[:index], posts[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(posts)
}

func createPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var post Post
	_ = json.NewDecoder(r.Body).Decode(post)
	post.ID = strconv.Itoa(rand.Intn(1000000))
	posts = append(posts, post)
	json.NewEncoder(w).Encode(&post)
}

func saveMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var message Message
	_ = json.NewDecoder(r.Body).Decode(message)
	// post.ID = strconv.Itoa(rand.Intn(1000000))
	// posts = append(posts, post)
	// json.NewEncoder(w).Encode(&message)
}

func getMessage(w http.ResponseWriter, r *http.Request) {
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

func setMessage(w http.ResponseWriter, r *http.Request) {

}

func main() {
	// conn, err := redis.Dial("tcp", "localhost:6379")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// defer conn.Close()

	// _, err = conn.Do("HMSET", "album:2", "title", "Electric Ladyland", "artist", "Jimi Hendrix", "price", 4.95, "likes", 8)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// title, err := redis.String(conn.Do("HGET", "album:1", "title"))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// artist, err := redis.String(conn.Do("HGET", "album:1", "artist"))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// price, err := redis.Float64(conn.Do("HGET", "album:1", "price"))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// likes, err := redis.Int(conn.Do("HGET", "album:1", "likes"))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("%s by %s: £%.2f [%d likes]\n", title, artist, price, likes)

	// fmt.Println("Electric Ladyland added!")

	pool = &redis.Pool{
		MaxIdle:     10,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379")
		},
	}

	router := mux.NewRouter()

	posts = append(posts, Post{ID: "1", Title: "My first post", Body: "This is the content of my first post"})

	router.HandleFunc("/posts", getPosts).Methods("GET")
	router.HandleFunc("/posts", createPost).Methods("POST")
	router.HandleFunc("/posts/{id}", getPost).Methods("GET")
	router.HandleFunc("/posts/{id}", updatePost).Methods("PUT")
	router.HandleFunc("/posts/{id}", deletePost).Methods("DELETE")

	router.HandleFunc("/message", saveMessage).Methods("POST")
	router.HandleFunc("/message", getMessage).Methods("GET")

	http.ListenAndServe(":8000", router)

}
