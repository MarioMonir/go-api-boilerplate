package server

import (
	"fmt"
	"go-api-boilerplate/handlers"
	"net/http"
	"time"
)

const (
	// Host name of the HTTP Server
	Host = "localhost"
	// Port of the HTTP Server
	Port = ":8080"
)

func newServer() *http.Server {
	// Create a multiplexer
	mux := http.NewServeMux()

	// Register handlers with the multiplexer
	mux.HandleFunc("/", handlers.RootHandler)
	mux.HandleFunc("/users", handlers.UsersHandler)

	// Create and return the server
	return &http.Server{
		Addr:           Port,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}

func StartWebServer() {
	srv := newServer()

	fmt.Println("Server is Listening at port", Port)

	if err := srv.ListenAndServe(); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
