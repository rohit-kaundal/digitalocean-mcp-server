package handlers

import (
	"context"
	"fmt"

	"github.com/digitalocean/godo"
	mcp_golang "github.com/metoro-io/mcp-golang"
)

func (h *Handler) ListFloatingIPs() (*mcp_golang.ToolResponse, error) {
	client := h.doClient.GetClient()
	
	floatingIPs, _, err := client.FloatingIPs.List(context.Background(), &godo.ListOptions{})
	if err != nil {
		return h.HandleError(err, "list_floating_ips")
	}

	return h.HandleSuccess(floatingIPs, "list_floating_ips")
}

func (h *Handler) GetFloatingIP(ip string) (*mcp_golang.ToolResponse, error) {
	client := h.doClient.GetClient()
	
	floatingIP, _, err := client.FloatingIPs.Get(context.Background(), ip)
	if err != nil {
		return h.HandleError(err, "get_floating_ip")
	}

	return h.HandleSuccess(floatingIP, "get_floating_ip")
}

func (h *Handler) CreateFloatingIP(region string, dropletID int) (*mcp_golang.ToolResponse, error) {
	client := h.doClient.GetClient()
	
	var createRequest *godo.FloatingIPCreateRequest
	
	if dropletID > 0 {
		// Assign to specific droplet
		createRequest = &godo.FloatingIPCreateRequest{
			DropletID: dropletID,
		}
	} else {
		// Create reserved floating IP for region
		createRequest = &godo.FloatingIPCreateRequest{
			Region: region,
		}
	}
	
	floatingIP, _, err := client.FloatingIPs.Create(context.Background(), createRequest)
	if err != nil {
		return h.HandleError(err, "create_floating_ip")
	}

	return h.HandleSuccess(floatingIP, "create_floating_ip")
}

func (h *Handler) DeleteFloatingIP(ip string) (*mcp_golang.ToolResponse, error) {
	client := h.doClient.GetClient()
	
	_, err := client.FloatingIPs.Delete(context.Background(), ip)
	if err != nil {
		return h.HandleError(err, "delete_floating_ip")
	}

	return h.HandleSuccess(map[string]string{
		"status":  "success",
		"message": fmt.Sprintf("Floating IP %s deleted successfully", ip),
	}, "delete_floating_ip")
}

func (h *Handler) AssignFloatingIP(ip string, dropletID int) (*mcp_golang.ToolResponse, error) {
	client := h.doClient.GetClient()
	
	action, _, err := client.FloatingIPActions.Assign(context.Background(), ip, dropletID)
	if err != nil {
		return h.HandleError(err, "assign_floating_ip")
	}

	return h.HandleSuccess(action, "assign_floating_ip")
}

func (h *Handler) UnassignFloatingIP(ip string) (*mcp_golang.ToolResponse, error) {
	client := h.doClient.GetClient()
	
	action, _, err := client.FloatingIPActions.Unassign(context.Background(), ip)
	if err != nil {
		return h.HandleError(err, "unassign_floating_ip")
	}

	return h.HandleSuccess(action, "unassign_floating_ip")
}