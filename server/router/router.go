package entangle

import (
	"../middleware"
	"github.com/gorilla/mux"
)

func Router() {

	router := mux.NewRouter()

	router.HandleFunc("/posts", middleware.GetPosts).Methods("GET")
	router.HandleFunc("/posts", middleware.createPost).Methods("POST")
	router.HandleFunc("/posts/{id}", middleware.getPost).Methods("GET")
	router.HandleFunc("/posts/{id}", middleware.updatePost).Methods("PUT")
	router.HandleFunc("/posts/{id}", middleware.deletePost).Methods("DELETE")

	router.HandleFunc("/message", middleware.saveMessage).Methods("POST")
	router.HandleFunc("/message", middleware.getMessage).Methods("GET")
}
