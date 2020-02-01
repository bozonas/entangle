package main

import (
	"fmt"
	"log"
	"net/http"

	"./router"
)

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

	r := router.Router()
	// fs := http.FileServer(http.Dir("build"))
	// http.Handle("/", fs)
	fmt.Println("Starting server on the port 8080...")

	log.Fatal(http.ListenAndServe(":8080", r))
}
