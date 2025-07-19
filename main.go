package main

import (
	"digitalocean-mcp-server/server"
	"log"
	"os"
)

func main() {
	done := make(chan struct{})

	if err := server.Start(); err != nil {
		log.Printf("Server failed to start: %v", err)
		os.Exit(1)
	}
	<-done

}
