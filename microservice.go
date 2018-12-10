package main

import (
	"net/http"
	"os"

	"sirfilip/cloud-native-go/api"
)

func main() {
	http.HandleFunc("/", api.Hello)
	http.HandleFunc("/api/echo", api.Echo)
	http.HandleFunc("/api/books", api.BooksHandlerFunc)
	http.HandleFunc("/api/books/", api.BookHandlerFunc)
	http.ListenAndServe(port(), nil)
}

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	return ":" + port
}
