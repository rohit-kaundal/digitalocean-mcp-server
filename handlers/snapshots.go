package handlers

import (
	"context"
	"fmt"

	"github.com/digitalocean/godo"
	mcp_golang "github.com/metoro-io/mcp-golang"
)

func (h *Handler) ListSnapshots(resourceType string) (*mcp_golang.ToolResponse, error) {
	client := h.doClient.GetClient()
	
	listOptions := &godo.ListOptions{}
	
	snapshots, _, err := client.Snapshots.List(context.Background(), listOptions)
	if err != nil {
		return h.HandleError(err, "list_snapshots")
	}

	// Filter by resource type if specified
	if resourceType != "" {
		var filteredSnapshots []godo.Snapshot
		for _, snapshot := range snapshots {
			if snapshot.ResourceType == resourceType {
				filteredSnapshots = append(filteredSnapshots, snapshot)
			}
		}
		return h.HandleSuccess(filteredSnapshots, "list_snapshots")
	}

	return h.HandleSuccess(snapshots, "list_snapshots")
}

func (h *Handler) ListVolumeSnapshots() (*mcp_golang.ToolResponse, error) {
	client := h.doClient.GetClient()
	
	snapshots, _, err := client.Snapshots.ListVolume(context.Background(), &godo.ListOptions{})
	if err != nil {
		return h.HandleError(err, "list_volume_snapshots")
	}

	return h.HandleSuccess(snapshots, "list_volume_snapshots")
}

func (h *Handler) ListDropletSnapshots() (*mcp_golang.ToolResponse, error) {
	client := h.doClient.GetClient()
	
	snapshots, _, err := client.Snapshots.ListDroplet(context.Background(), &godo.ListOptions{})
	if err != nil {
		return h.HandleError(err, "list_droplet_snapshots")
	}

	return h.HandleSuccess(snapshots, "list_droplet_snapshots")
}

func (h *Handler) GetSnapshot(snapshotID string) (*mcp_golang.ToolResponse, error) {
	client := h.doClient.GetClient()
	
	snapshot, _, err := client.Snapshots.Get(context.Background(), snapshotID)
	if err != nil {
		return h.HandleError(err, "get_snapshot")
	}

	return h.HandleSuccess(snapshot, "get_snapshot")
}

func (h *Handler) DeleteSnapshot(snapshotID string) (*mcp_golang.ToolResponse, error) {
	client := h.doClient.GetClient()
	
	_, err := client.Snapshots.Delete(context.Background(), snapshotID)
	if err != nil {
		return h.HandleError(err, "delete_snapshot")
	}

	return h.HandleSuccess(map[string]string{
		"status":  "success",
		"message": fmt.Sprintf("Snapshot %s deleted successfully", snapshotID),
	}, "delete_snapshot")
}

func (h *Handler) CreateDropletSnapshot(dropletID int, name string) (*mcp_golang.ToolResponse, error) {
	client := h.doClient.GetClient()
	
	action, _, err := client.DropletActions.Snapshot(context.Background(), dropletID, name)
	if err != nil {
		return h.HandleError(err, "create_droplet_snapshot")
	}

	return h.HandleSuccess(action, "create_droplet_snapshot")
}