package handlers

import (
	"context"
	"fmt"

	"github.com/digitalocean/godo"
	mcp_golang "github.com/metoro-io/mcp-golang"
)

func (h *Handler) ListVolumes(region string) (*mcp_golang.ToolResponse, error) {
	client := h.doClient.GetClient()
	
	listOptions := &godo.ListVolumeParams{}
	if region != "" {
		listOptions.Region = region
	}
	
	volumes, _, err := client.Storage.ListVolumes(context.Background(), listOptions)
	if err != nil {
		return h.HandleError(err, "list_volumes")
	}

	return h.HandleSuccess(volumes, "list_volumes")
}

func (h *Handler) GetVolume(volumeID string) (*mcp_golang.ToolResponse, error) {
	client := h.doClient.GetClient()
	
	volume, _, err := client.Storage.GetVolume(context.Background(), volumeID)
	if err != nil {
		return h.HandleError(err, "get_volume")
	}

	return h.HandleSuccess(volume, "get_volume")
}

func (h *Handler) CreateVolume(name, region string, sizeGigaBytes int64, description string) (*mcp_golang.ToolResponse, error) {
	client := h.doClient.GetClient()
	
	createRequest := &godo.VolumeCreateRequest{
		Name:          name,
		Region:        region,
		SizeGigaBytes: sizeGigaBytes,
		Description:   description,
	}
	
	volume, _, err := client.Storage.CreateVolume(context.Background(), createRequest)
	if err != nil {
		return h.HandleError(err, "create_volume")
	}

	return h.HandleSuccess(volume, "create_volume")
}

func (h *Handler) DeleteVolume(volumeID string) (*mcp_golang.ToolResponse, error) {
	client := h.doClient.GetClient()
	
	_, err := client.Storage.DeleteVolume(context.Background(), volumeID)
	if err != nil {
		return h.HandleError(err, "delete_volume")
	}

	return h.HandleSuccess(map[string]string{
		"status":  "success",
		"message": fmt.Sprintf("Volume %s deleted successfully", volumeID),
	}, "delete_volume")
}

func (h *Handler) AttachVolume(volumeID string, dropletID int) (*mcp_golang.ToolResponse, error) {
	client := h.doClient.GetClient()
	
	action, _, err := client.StorageActions.Attach(context.Background(), volumeID, dropletID)
	if err != nil {
		return h.HandleError(err, "attach_volume")
	}

	return h.HandleSuccess(action, "attach_volume")
}

func (h *Handler) DetachVolume(volumeID string, dropletID int) (*mcp_golang.ToolResponse, error) {
	client := h.doClient.GetClient()
	
	action, _, err := client.StorageActions.DetachByDropletID(context.Background(), volumeID, dropletID)
	if err != nil {
		return h.HandleError(err, "detach_volume")
	}

	return h.HandleSuccess(action, "detach_volume")
}

func (h *Handler) ResizeVolume(volumeID string, sizeGigaBytes int64, region string) (*mcp_golang.ToolResponse, error) {
	client := h.doClient.GetClient()
	
	action, _, err := client.StorageActions.Resize(context.Background(), volumeID, int(sizeGigaBytes), region)
	if err != nil {
		return h.HandleError(err, "resize_volume")
	}

	return h.HandleSuccess(action, "resize_volume")
}

func (h *Handler) CreateVolumeSnapshot(volumeID, name, description string) (*mcp_golang.ToolResponse, error) {
	client := h.doClient.GetClient()
	
	createRequest := &godo.SnapshotCreateRequest{
		VolumeID:    volumeID,
		Name:        name,
		Description: description,
	}
	
	snapshot, _, err := client.Storage.CreateSnapshot(context.Background(), createRequest)
	if err != nil {
		return h.HandleError(err, "create_volume_snapshot")
	}

	return h.HandleSuccess(snapshot, "create_volume_snapshot")
}