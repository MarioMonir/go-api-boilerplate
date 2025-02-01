package server

import (
	"fmt"
	"go-api-boilerplate/modules/user/user_handlers"
	"net/http"
	"time"
)

const (
	// Host name of the HTTP Server
	Host = "localhost"
	// Port of the HTTP Server
	Port = ":8080"
)

func LaunchWebServer() {
	// Create a multiplexer
	mux := http.NewServeMux()

	userHandler := user_handlers.NewUserHandler()

	// Register handlers with the multiplexer
	userHandler.RegisterUserHanlders(mux)

	// Create and return the server
	srv := &http.Server{
		Addr:           Port,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	fmt.Println("Server is Listening at port", Port)

	if err := srv.ListenAndServe(); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
