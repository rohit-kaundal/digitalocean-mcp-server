package handlers

import (
	"context"
	"fmt"

	"github.com/digitalocean/godo"
	mcp_golang "github.com/metoro-io/mcp-golang"
)

func (h *Handler) ListDroplets(page, perPage int) (*mcp_golang.ToolResponse, error) {
	client := h.doClient.GetClient()
	
	// Set default values if not provided
	if page <= 0 {
		page = 1
	}
	if perPage <= 0 {
		perPage = 25
	}
	// Limit perPage to DigitalOcean's maximum of 200
	if perPage > 200 {
		perPage = 200
	}
	
	droplets, response, err := client.Droplets.List(context.Background(), &godo.ListOptions{
		Page:    page,
		PerPage: perPage,
	})
	if err != nil {
		return h.HandleError(err, "list_droplets")
	}

	// Transform droplets to only include ID, Name, and Status
	simplifiedDroplets := make([]map[string]interface{}, len(droplets))
	for i, droplet := range droplets {
		simplifiedDroplets[i] = map[string]interface{}{
			"id":     droplet.ID,
			"name":   droplet.Name,
			"status": droplet.Status,
		}
	}

	// Include pagination metadata in response
	result := map[string]interface{}{
		"droplets": simplifiedDroplets,
		"meta": map[string]interface{}{
			"total":    response.Meta.Total,
			"page":     page,
			"per_page": perPage,
			"pages":    (response.Meta.Total + perPage - 1) / perPage,
		},
		"links": response.Links,
	}

	return h.HandleSuccess(result, "list_droplets")
}

func (h *Handler) GetDroplet(dropletID int) (*mcp_golang.ToolResponse, error) {
	client := h.doClient.GetClient()
	
	droplet, _, err := client.Droplets.Get(context.Background(), dropletID)
	if err != nil {
		return h.HandleError(err, "get_droplet")
	}

	return h.HandleSuccess(droplet, "get_droplet")
}

func (h *Handler) CreateDroplet(name, region, size, image string) (*mcp_golang.ToolResponse, error) {
	client := h.doClient.GetClient()
	
	createRequest := &godo.DropletCreateRequest{
		Name:   name,
		Region: region,
		Size:   size,
		Image: godo.DropletCreateImage{
			Slug: image,
		},
	}
	
	droplet, _, err := client.Droplets.Create(context.Background(), createRequest)
	if err != nil {
		return h.HandleError(err, "create_droplet")
	}

	return h.HandleSuccess(droplet, "create_droplet")
}

func (h *Handler) DeleteDroplet(dropletID int) (*mcp_golang.ToolResponse, error) {
	client := h.doClient.GetClient()
	
	_, err := client.Droplets.Delete(context.Background(), dropletID)
	if err != nil {
		return h.HandleError(err, "delete_droplet")
	}

	return h.HandleSuccess(map[string]string{
		"status":  "success",
		"message": fmt.Sprintf("Droplet %d deleted successfully", dropletID),
	}, "delete_droplet")
}

func (h *Handler) ResizeDroplet(dropletID int, size string, disk bool) (*mcp_golang.ToolResponse, error) {
	client := h.doClient.GetClient()
	
	_, _, err := client.DropletActions.Resize(context.Background(), dropletID, size, disk)
	if err != nil {
		return h.HandleError(err, "resize_droplet")
	}

	return h.HandleSuccess(map[string]string{
		"status":  "success",
		"message": fmt.Sprintf("Droplet %d resize initiated", dropletID),
	}, "resize_droplet")
}