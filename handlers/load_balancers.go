package handlers

import (
	"context"
	"fmt"

	"github.com/digitalocean/godo"
	mcp_golang "github.com/metoro-io/mcp-golang"
)

func (h *Handler) ListLoadBalancers() (*mcp_golang.ToolResponse, error) {
	client := h.doClient.GetClient()
	
	loadBalancers, _, err := client.LoadBalancers.List(context.Background(), &godo.ListOptions{})
	if err != nil {
		return h.HandleError(err, "list_load_balancers")
	}

	return h.HandleSuccess(loadBalancers, "list_load_balancers")
}

func (h *Handler) GetLoadBalancer(lbID string) (*mcp_golang.ToolResponse, error) {
	client := h.doClient.GetClient()
	
	loadBalancer, _, err := client.LoadBalancers.Get(context.Background(), lbID)
	if err != nil {
		return h.HandleError(err, "get_load_balancer")
	}

	return h.HandleSuccess(loadBalancer, "get_load_balancer")
}

func (h *Handler) CreateLoadBalancer(name, algorithm, region string, forwardingRules []godo.ForwardingRule, dropletIDs []int) (*mcp_golang.ToolResponse, error) {
	client := h.doClient.GetClient()
	
	createRequest := &godo.LoadBalancerRequest{
		Name:            name,
		Algorithm:       algorithm,
		Region:          region,
		ForwardingRules: forwardingRules,
		DropletIDs:      dropletIDs,
		RedirectHttpToHttps: false,
		EnableProxyProtocol: false,
	}
	
	loadBalancer, _, err := client.LoadBalancers.Create(context.Background(), createRequest)
	if err != nil {
		return h.HandleError(err, "create_load_balancer")
	}

	return h.HandleSuccess(loadBalancer, "create_load_balancer")
}

func (h *Handler) UpdateLoadBalancer(lbID, name, algorithm, region string, forwardingRules []godo.ForwardingRule, dropletIDs []int) (*mcp_golang.ToolResponse, error) {
	client := h.doClient.GetClient()
	
	updateRequest := &godo.LoadBalancerRequest{
		Name:            name,
		Algorithm:       algorithm,
		Region:          region,
		ForwardingRules: forwardingRules,
		DropletIDs:      dropletIDs,
	}
	
	loadBalancer, _, err := client.LoadBalancers.Update(context.Background(), lbID, updateRequest)
	if err != nil {
		return h.HandleError(err, "update_load_balancer")
	}

	return h.HandleSuccess(loadBalancer, "update_load_balancer")
}

func (h *Handler) DeleteLoadBalancer(lbID string) (*mcp_golang.ToolResponse, error) {
	client := h.doClient.GetClient()
	
	_, err := client.LoadBalancers.Delete(context.Background(), lbID)
	if err != nil {
		return h.HandleError(err, "delete_load_balancer")
	}

	return h.HandleSuccess(map[string]string{
		"status":  "success",
		"message": fmt.Sprintf("Load balancer %s deleted successfully", lbID),
	}, "delete_load_balancer")
}

func (h *Handler) AddDropletsToLoadBalancer(lbID string, dropletIDs []int) (*mcp_golang.ToolResponse, error) {
	client := h.doClient.GetClient()
	
	_, err := client.LoadBalancers.AddDroplets(context.Background(), lbID, dropletIDs...)
	if err != nil {
		return h.HandleError(err, "add_droplets_to_load_balancer")
	}

	return h.HandleSuccess(map[string]string{
		"status":  "success",
		"message": fmt.Sprintf("Droplets added to load balancer %s successfully", lbID),
	}, "add_droplets_to_load_balancer")
}

func (h *Handler) RemoveDropletsFromLoadBalancer(lbID string, dropletIDs []int) (*mcp_golang.ToolResponse, error) {
	client := h.doClient.GetClient()
	
	_, err := client.LoadBalancers.RemoveDroplets(context.Background(), lbID, dropletIDs...)
	if err != nil {
		return h.HandleError(err, "remove_droplets_from_load_balancer")
	}

	return h.HandleSuccess(map[string]string{
		"status":  "success",
		"message": fmt.Sprintf("Droplets removed from load balancer %s successfully", lbID),
	}, "remove_droplets_from_load_balancer")
}

func (h *Handler) AddForwardingRulesToLoadBalancer(lbID string, forwardingRules []godo.ForwardingRule) (*mcp_golang.ToolResponse, error) {
	client := h.doClient.GetClient()
	
	_, err := client.LoadBalancers.AddForwardingRules(context.Background(), lbID, forwardingRules...)
	if err != nil {
		return h.HandleError(err, "add_forwarding_rules_to_load_balancer")
	}

	return h.HandleSuccess(map[string]string{
		"status":  "success",
		"message": fmt.Sprintf("Forwarding rules added to load balancer %s successfully", lbID),
	}, "add_forwarding_rules_to_load_balancer")
}

func (h *Handler) RemoveForwardingRulesFromLoadBalancer(lbID string, forwardingRules []godo.ForwardingRule) (*mcp_golang.ToolResponse, error) {
	client := h.doClient.GetClient()
	
	_, err := client.LoadBalancers.RemoveForwardingRules(context.Background(), lbID, forwardingRules...)
	if err != nil {
		return h.HandleError(err, "remove_forwarding_rules_from_load_balancer")
	}

	return h.HandleSuccess(map[string]string{
		"status":  "success",
		"message": fmt.Sprintf("Forwarding rules removed from load balancer %s successfully", lbID),
	}, "remove_forwarding_rules_from_load_balancer")
}