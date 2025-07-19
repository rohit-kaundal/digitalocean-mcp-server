package handlers

import (
	"context"
	"fmt"

	"github.com/digitalocean/godo"
	mcp_golang "github.com/metoro-io/mcp-golang"
)

func (h *Handler) ListK8SClusters() (*mcp_golang.ToolResponse, error) {
	client := h.doClient.GetClient()
	
	clusters, _, err := client.Kubernetes.List(context.Background(), &godo.ListOptions{})
	if err != nil {
		return h.HandleError(err, "list_k8s_clusters")
	}

	return h.HandleSuccess(clusters, "list_k8s_clusters")
}

func (h *Handler) GetK8SCluster(clusterID string) (*mcp_golang.ToolResponse, error) {
	client := h.doClient.GetClient()
	
	cluster, _, err := client.Kubernetes.Get(context.Background(), clusterID)
	if err != nil {
		return h.HandleError(err, "get_k8s_cluster")
	}

	return h.HandleSuccess(cluster, "get_k8s_cluster")
}

func (h *Handler) CreateK8SCluster(name, region, version, nodePoolSize string, nodeCount int) (*mcp_golang.ToolResponse, error) {
	client := h.doClient.GetClient()
	
	createRequest := &godo.KubernetesClusterCreateRequest{
		Name:        name,
		RegionSlug:  region,
		VersionSlug: version,
		NodePools: []*godo.KubernetesNodePoolCreateRequest{
			{
				Name:  fmt.Sprintf("%s-pool", name),
				Size:  nodePoolSize,
				Count: nodeCount,
			},
		},
	}
	
	cluster, _, err := client.Kubernetes.Create(context.Background(), createRequest)
	if err != nil {
		return h.HandleError(err, "create_k8s_cluster")
	}

	return h.HandleSuccess(cluster, "create_k8s_cluster")
}

func (h *Handler) DeleteK8SCluster(clusterID string) (*mcp_golang.ToolResponse, error) {
	client := h.doClient.GetClient()
	
	_, err := client.Kubernetes.Delete(context.Background(), clusterID)
	if err != nil {
		return h.HandleError(err, "delete_k8s_cluster")
	}

	return h.HandleSuccess(map[string]string{
		"status":  "success",
		"message": fmt.Sprintf("Kubernetes cluster %s deleted successfully", clusterID),
	}, "delete_k8s_cluster")
}

func (h *Handler) GetK8SClusterKubeconfig(clusterID string) (*mcp_golang.ToolResponse, error) {
	client := h.doClient.GetClient()
	
	kubeconfig, _, err := client.Kubernetes.GetKubeConfig(context.Background(), clusterID)
	if err != nil {
		return h.HandleError(err, "get_k8s_cluster_kubeconfig")
	}

	return h.HandleSuccess(map[string]string{
		"kubeconfig": string(kubeconfig.KubeconfigYAML),
	}, "get_k8s_cluster_kubeconfig")
}

func (h *Handler) ListK8SNodePools(clusterID string) (*mcp_golang.ToolResponse, error) {
	client := h.doClient.GetClient()
	
	nodePools, _, err := client.Kubernetes.ListNodePools(context.Background(), clusterID, &godo.ListOptions{})
	if err != nil {
		return h.HandleError(err, "list_k8s_node_pools")
	}

	return h.HandleSuccess(nodePools, "list_k8s_node_pools")
}

func (h *Handler) GetK8SNodePool(clusterID, poolID string) (*mcp_golang.ToolResponse, error) {
	client := h.doClient.GetClient()
	
	nodePool, _, err := client.Kubernetes.GetNodePool(context.Background(), clusterID, poolID)
	if err != nil {
		return h.HandleError(err, "get_k8s_node_pool")
	}

	return h.HandleSuccess(nodePool, "get_k8s_node_pool")
}