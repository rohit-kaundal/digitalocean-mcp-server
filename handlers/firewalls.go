package handlers

import (
	"context"
	"fmt"

	"github.com/digitalocean/godo"
	mcp_golang "github.com/metoro-io/mcp-golang"
)

func (h *Handler) ListFirewalls() (*mcp_golang.ToolResponse, error) {
	client := h.doClient.GetClient()
	
	firewalls, _, err := client.Firewalls.List(context.Background(), &godo.ListOptions{})
	if err != nil {
		return h.HandleError(err, "list_firewalls")
	}

	simplifiedFirewalls := make([]map[string]interface{}, len(firewalls))
	for i, firewall := range firewalls {
		simplifiedFirewalls[i] = map[string]interface{}{
			"id":     firewall.ID,
			"name":   firewall.Name,
			"status": firewall.Status,
		}
	}

	result := map[string]interface{}{
		"firewalls": simplifiedFirewalls,
	}

	return h.HandleSuccess(result, "list_firewalls")
}

func (h *Handler) GetFirewall(firewallID string) (*mcp_golang.ToolResponse, error) {
	client := h.doClient.GetClient()
	
	firewall, _, err := client.Firewalls.Get(context.Background(), firewallID)
	if err != nil {
		return h.HandleError(err, "get_firewall")
	}

	return h.HandleSuccess(firewall, "get_firewall")
}

func (h *Handler) CreateFirewall(name string, inboundRules []godo.InboundRule, outboundRules []godo.OutboundRule, dropletIDs []int, tags []string) (*mcp_golang.ToolResponse, error) {
	client := h.doClient.GetClient()
	
	createRequest := &godo.FirewallRequest{
		Name:          name,
		InboundRules:  inboundRules,
		OutboundRules: outboundRules,
		DropletIDs:    dropletIDs,
		Tags:          tags,
	}
	
	firewall, _, err := client.Firewalls.Create(context.Background(), createRequest)
	if err != nil {
		return h.HandleError(err, "create_firewall")
	}

	return h.HandleSuccess(firewall, "create_firewall")
}

func (h *Handler) UpdateFirewall(firewallID, name string, inboundRules []godo.InboundRule, outboundRules []godo.OutboundRule) (*mcp_golang.ToolResponse, error) {
	client := h.doClient.GetClient()
	
	updateRequest := &godo.FirewallRequest{
		Name:          name,
		InboundRules:  inboundRules,
		OutboundRules: outboundRules,
	}
	
	firewall, _, err := client.Firewalls.Update(context.Background(), firewallID, updateRequest)
	if err != nil {
		return h.HandleError(err, "update_firewall")
	}

	return h.HandleSuccess(firewall, "update_firewall")
}

func (h *Handler) DeleteFirewall(firewallID string) (*mcp_golang.ToolResponse, error) {
	client := h.doClient.GetClient()
	
	_, err := client.Firewalls.Delete(context.Background(), firewallID)
	if err != nil {
		return h.HandleError(err, "delete_firewall")
	}

	return h.HandleSuccess(map[string]string{
		"status":  "success",
		"message": fmt.Sprintf("Firewall %s deleted successfully", firewallID),
	}, "delete_firewall")
}

func (h *Handler) AddDropletsToFirewall(firewallID string, dropletIDs []int) (*mcp_golang.ToolResponse, error) {
	client := h.doClient.GetClient()
	
	_, err := client.Firewalls.AddDroplets(context.Background(), firewallID, dropletIDs...)
	if err != nil {
		return h.HandleError(err, "add_droplets_to_firewall")
	}

	return h.HandleSuccess(map[string]string{
		"status":  "success",
		"message": fmt.Sprintf("Droplets added to firewall %s successfully", firewallID),
	}, "add_droplets_to_firewall")
}

func (h *Handler) RemoveDropletsFromFirewall(firewallID string, dropletIDs []int) (*mcp_golang.ToolResponse, error) {
	client := h.doClient.GetClient()
	
	_, err := client.Firewalls.RemoveDroplets(context.Background(), firewallID, dropletIDs...)
	if err != nil {
		return h.HandleError(err, "remove_droplets_from_firewall")
	}

	return h.HandleSuccess(map[string]string{
		"status":  "success",
		"message": fmt.Sprintf("Droplets removed from firewall %s successfully", firewallID),
	}, "remove_droplets_from_firewall")
}

func (h *Handler) AddTagsToFirewall(firewallID string, tags []string) (*mcp_golang.ToolResponse, error) {
	client := h.doClient.GetClient()
	
	_, err := client.Firewalls.AddTags(context.Background(), firewallID, tags...)
	if err != nil {
		return h.HandleError(err, "add_tags_to_firewall")
	}

	return h.HandleSuccess(map[string]string{
		"status":  "success",
		"message": fmt.Sprintf("Tags added to firewall %s successfully", firewallID),
	}, "add_tags_to_firewall")
}

func (h *Handler) RemoveTagsFromFirewall(firewallID string, tags []string) (*mcp_golang.ToolResponse, error) {
	client := h.doClient.GetClient()
	
	_, err := client.Firewalls.RemoveTags(context.Background(), firewallID, tags...)
	if err != nil {
		return h.HandleError(err, "remove_tags_from_firewall")
	}

	return h.HandleSuccess(map[string]string{
		"status":  "success",
		"message": fmt.Sprintf("Tags removed from firewall %s successfully", firewallID),
	}, "remove_tags_from_firewall")
}

func (h *Handler) AddRulesToFirewall(firewallID string, inboundRules []godo.InboundRule, outboundRules []godo.OutboundRule) (*mcp_golang.ToolResponse, error) {
	client := h.doClient.GetClient()
	
	rulesRequest := &godo.FirewallRulesRequest{
		InboundRules:  inboundRules,
		OutboundRules: outboundRules,
	}
	
	_, err := client.Firewalls.AddRules(context.Background(), firewallID, rulesRequest)
	if err != nil {
		return h.HandleError(err, "add_rules_to_firewall")
	}

	return h.HandleSuccess(map[string]string{
		"status":  "success",
		"message": fmt.Sprintf("Rules added to firewall %s successfully", firewallID),
	}, "add_rules_to_firewall")
}

func (h *Handler) RemoveRulesFromFirewall(firewallID string, inboundRules []godo.InboundRule, outboundRules []godo.OutboundRule) (*mcp_golang.ToolResponse, error) {
	client := h.doClient.GetClient()
	
	rulesRequest := &godo.FirewallRulesRequest{
		InboundRules:  inboundRules,
		OutboundRules: outboundRules,
	}
	
	_, err := client.Firewalls.RemoveRules(context.Background(), firewallID, rulesRequest)
	if err != nil {
		return h.HandleError(err, "remove_rules_from_firewall")
	}

	return h.HandleSuccess(map[string]string{
		"status":  "success",
		"message": fmt.Sprintf("Rules removed from firewall %s successfully", firewallID),
	}, "remove_rules_from_firewall")
}