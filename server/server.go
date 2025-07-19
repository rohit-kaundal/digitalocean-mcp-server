package server

import (
	"digitalocean-mcp-server/client"
	"digitalocean-mcp-server/handlers"
	"log"

	mcp_golang "github.com/metoro-io/mcp-golang"
	"github.com/metoro-io/mcp-golang/transport/stdio"
)

func NewServer() (*mcp_golang.Server, error) {
	doClient, err := client.NewDOClient()
	if err != nil {
		return nil, err
	}

	handler := handlers.NewHandler(doClient)
	server := mcp_golang.NewServer(stdio.NewStdioServerTransport())

	if err := RegisterTools(server, handler); err != nil {
		return nil, err
	}

	return server, nil
}

func Start() error {
	server, err := NewServer()
	if err != nil {
		log.Fatalf("Failed to create server: %v", err)
		return err
	}

	log.Println("Starting DigitalOcean MCP server...")
	if err := server.Serve(); err != nil {
		log.Fatalf("Server error: %v", err)
		return err
	}

	return nil
}