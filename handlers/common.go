package handlers

import (
	"digitalocean-mcp-server/client"
	"encoding/json"
	"fmt"

	mcp_golang "github.com/metoro-io/mcp-golang"
)

type Handler struct {
	doClient *client.DOClient
}

func NewHandler(doClient *client.DOClient) *Handler {
	return &Handler{
		doClient: doClient,
	}
}

func (h *Handler) HandleError(err error, operation string) (*mcp_golang.ToolResponse, error) {
	return nil, fmt.Errorf("error in %s: %v", operation, err)
}

func (h *Handler) HandleSuccess(data interface{}, operation string) (*mcp_golang.ToolResponse, error) {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return h.HandleError(err, operation+" (JSON marshaling)")
	}

	return mcp_golang.NewToolResponse(
		mcp_golang.NewTextContent(string(jsonData)),
	), nil
}

func (h *Handler) GetDOClient() *client.DOClient {
	return h.doClient
}

func (h *Handler) TestConnection() (*mcp_golang.ToolResponse, error) {
	err := h.doClient.TestConnection()
	if err != nil {
		return h.HandleError(err, "connection test")
	}
	
	return h.HandleSuccess(map[string]string{
		"status": "connected",
		"message": "Successfully connected to DigitalOcean API",
	}, "connection test")
}