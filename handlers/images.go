package handlers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/digitalocean/godo"
	mcp_golang "github.com/metoro-io/mcp-golang"
)

func (h *Handler) ListImages(imageType string, isPublic bool) (*mcp_golang.ToolResponse, error) {
	client := h.doClient.GetClient()
	
	listOptions := &godo.ListOptions{}
	
	var images []godo.Image
	var err error
	
	switch imageType {
	case "distribution":
		images, _, err = client.Images.ListDistribution(context.Background(), listOptions)
	case "application":
		images, _, err = client.Images.ListApplication(context.Background(), listOptions)
	case "user":
		images, _, err = client.Images.ListUser(context.Background(), listOptions)
	default:
		// List all images
		images, _, err = client.Images.List(context.Background(), listOptions)
	}
	
	if err != nil {
		return h.HandleError(err, "list_images")
	}

	return h.HandleSuccess(images, "list_images")
}

func (h *Handler) GetImage(imageID string) (*mcp_golang.ToolResponse, error) {
	client := h.doClient.GetClient()
	
	// Try to parse as int ID first, then by slug if parsing fails
	if id, err := strconv.Atoi(imageID); err == nil {
		image, _, err := client.Images.GetByID(context.Background(), id)
		if err == nil {
			return h.HandleSuccess(image, "get_image")
		}
	}
	
	// Try by slug
	image, _, err := client.Images.GetBySlug(context.Background(), imageID)
	if err != nil {
		return h.HandleError(err, "get_image")
	}

	return h.HandleSuccess(image, "get_image")
}

func (h *Handler) UpdateImage(imageID, name string) (*mcp_golang.ToolResponse, error) {
	client := h.doClient.GetClient()
	
	id, err := strconv.Atoi(imageID)
	if err != nil {
		return h.HandleError(fmt.Errorf("invalid image ID: %s", imageID), "update_image")
	}
	
	updateRequest := &godo.ImageUpdateRequest{
		Name: name,
	}
	
	image, _, err := client.Images.Update(context.Background(), id, updateRequest)
	if err != nil {
		return h.HandleError(err, "update_image")
	}

	return h.HandleSuccess(image, "update_image")
}

func (h *Handler) DeleteImage(imageID string) (*mcp_golang.ToolResponse, error) {
	client := h.doClient.GetClient()
	
	id, err := strconv.Atoi(imageID)
	if err != nil {
		return h.HandleError(fmt.Errorf("invalid image ID: %s", imageID), "delete_image")
	}
	
	_, err = client.Images.Delete(context.Background(), id)
	if err != nil {
		return h.HandleError(err, "delete_image")
	}

	return h.HandleSuccess(map[string]string{
		"status":  "success",
		"message": fmt.Sprintf("Image %s deleted successfully", imageID),
	}, "delete_image")
}

func (h *Handler) TransferImage(imageID, regionSlug string) (*mcp_golang.ToolResponse, error) {
	client := h.doClient.GetClient()
	
	id, err := strconv.Atoi(imageID)
	if err != nil {
		return h.HandleError(fmt.Errorf("invalid image ID: %s", imageID), "transfer_image")
	}
	
	transferRequest := &godo.ActionRequest{
		"type":   "transfer",
		"region": regionSlug,
	}
	
	action, _, err := client.ImageActions.Transfer(context.Background(), id, transferRequest)
	if err != nil {
		return h.HandleError(err, "transfer_image")
	}

	return h.HandleSuccess(action, "transfer_image")
}

func (h *Handler) ConvertImageToSnapshot(imageID string) (*mcp_golang.ToolResponse, error) {
	client := h.doClient.GetClient()
	
	id, err := strconv.Atoi(imageID)
	if err != nil {
		return h.HandleError(fmt.Errorf("invalid image ID: %s", imageID), "convert_image_to_snapshot")
	}
	
	action, _, err := client.ImageActions.Convert(context.Background(), id)
	if err != nil {
		return h.HandleError(err, "convert_image_to_snapshot")
	}

	return h.HandleSuccess(action, "convert_image_to_snapshot")
}