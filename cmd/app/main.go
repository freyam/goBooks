package main

import (
	"fmt"
	"goBooks/app/app"
	"goBooks/app/router"
	"goBooks/config"
	lr "goBooks/util/logger"
	"net/http"
)

func main() {
	appConf := config.AppConfig()
	logger := lr.New(appConf.Debug)
	application := app.New(logger)
	appRouter := router.New(application)

	address := fmt.Sprintf(":%d", appConf.Server.Port)

	logger.Info().Msgf("Starting server %s\n", address)

	s := &http.Server{
		Addr:         address,
		Handler:      appRouter,
		ReadTimeout:  appConf.Server.TimeoutRead,
		WriteTimeout: appConf.Server.TimeoutWrite,
		IdleTimeout:  appConf.Server.TimeoutIdle,
	}
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatal().Err(err).Msg("Server startup failed")
	}
}

func Greet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}
