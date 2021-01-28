package main

import (
	"basic_chat_server/server"
	"log"
	"time"
)

func main() {
	var s = server.Server{
		Addr:        ":6666",
		IdleTimeout: 20 * time.Second,
	}
	err := s.ListenAndServe()
	if err != nil {
		log.Print(err)
	}
}
