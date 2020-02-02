package router

import (
	"net/http"

	"../middleware"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	buildHandler := http.FileServer(http.Dir("./client/build"))
	staticHandler := http.StripPrefix("/static/", http.FileServer(http.Dir("./client/build/static")))

	router := mux.NewRouter()

	router.HandleFunc("/api/message", middleware.SaveMessage).Methods("POST")
	router.HandleFunc("/api/message/{key}", middleware.GetMessage).Methods("GET")

	router.PathPrefix("/").Handler(buildHandler)
	router.PathPrefix("/static/").Handler(staticHandler)

	return router
}
