package main

import (
	"fmt"
	"goBooks/app/router"
	"goBooks/config"
	"log"
	"net/http"
)

func main() {
	appConf := config.AppConfig()
	appRouter := router.New()
	mux := http.NewServeMux()
	mux.HandleFunc("/", Greet)
	address := fmt.Sprintf(":%d", appConf.Server.Port)
	log.Printf("Starting server %s\n", address)
	s := &http.Server{
		Addr:         address,
		Handler:      appRouter,
		ReadTimeout:  appConf.Server.TimeoutRead,
		WriteTimeout: appConf.Server.TimeoutWrite,
		IdleTimeout:  appConf.Server.TimeoutIdle,
	}
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("Server startup failed")
	}
}

func Greet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}
