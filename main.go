package main

import (
	"github.com/daoraimi/dagger/server"
	"log"
)

func main() {
	s := server.New()
	if err := s.Run(); err != nil {
		log.Fatal("Failed to run service")
	}
}
