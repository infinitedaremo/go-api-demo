package main

import (
	"go-api-demo/internal/server"
	"go.uber.org/zap"
)

func main() {
	log, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	s, err := server.NewServer(log)
	if err != nil {
		log.Fatal("failed creating server", zap.Error(err))
	}

	log.Info("launching server")
	if err := s.Serve(":8080"); err != nil {
		log.Fatal("failed running server", zap.Error(err))
	}
}
