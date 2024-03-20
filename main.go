package main

import (
	"fmt"
	"log"
	"net/http"
)

func handleOther(w http.ResponseWriter, r *http.Request) {
	log.Println("Received a non domain request")
	w.Write([]byte("Hello, stranger..."))
}

func handle(w http.ResponseWriter, r *http.Request) {
	log.Println("Received a request at my domain")
	w.Write([]byte("Hello, Domain name!"))
}

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/", handleOther)
	router.HandleFunc("dreamsofcode.foo/", handle)

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	fmt.Println("Server listening on port :8080")
	server.ListenAndServe()
}
