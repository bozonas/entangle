package router

import (
	"net/http"

	"../middleware"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func Router() http.Handler {

	buildHandler := http.FileServer(http.Dir("./client/build"))
	staticHandler := http.StripPrefix("/static/", http.FileServer(http.Dir("./client/build/static")))

	router := mux.NewRouter()

	router.HandleFunc("/api/message", middleware.SaveMessage).Methods("POST")
	router.HandleFunc("/api/message/{key}", middleware.GetMessage).Methods("GET")

	router.PathPrefix("/").Handler(buildHandler)
	router.PathPrefix("/static/").Handler(staticHandler)

	cors := handlers.CORS(
		handlers.AllowedHeaders([]string{"content-type"}),
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowCredentials(),
	)

	return cors(router)
}
