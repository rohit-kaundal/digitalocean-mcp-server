package handlers

import (
	"context"
	"fmt"

	"github.com/digitalocean/godo"
	mcp_golang "github.com/metoro-io/mcp-golang"
)

func (h *Handler) ListRegistries() (*mcp_golang.ToolResponse, error) {
	client := h.doClient.GetClient()
	
	registry, _, err := client.Registry.Get(context.Background())
	if err != nil {
		return h.HandleError(err, "list_registries")
	}

	return h.HandleSuccess(registry, "list_registries")
}

func (h *Handler) GetRegistry(registryName string) (*mcp_golang.ToolResponse, error) {
	client := h.doClient.GetClient()
	
	registry, _, err := client.Registry.Get(context.Background())
	if err != nil {
		return h.HandleError(err, "get_registry")
	}

	return h.HandleSuccess(registry, "get_registry")
}

func (h *Handler) ListRepositories(registryName string) (*mcp_golang.ToolResponse, error) {
	client := h.doClient.GetClient()
	
	repositories, _, err := client.Registry.ListRepositories(context.Background(), registryName, &godo.ListOptions{})
	if err != nil {
		return h.HandleError(err, "list_repositories")
	}

	return h.HandleSuccess(repositories, "list_repositories")
}

func (h *Handler) GetRepository(registryName, repositoryName string) (*mcp_golang.ToolResponse, error) {
	client := h.doClient.GetClient()
	
	repositories, _, err := client.Registry.ListRepositories(context.Background(), registryName, &godo.ListOptions{})
	if err != nil {
		return h.HandleError(err, "get_repository")
	}

	for _, repo := range repositories {
		if repo.Name == repositoryName {
			return h.HandleSuccess(repo, "get_repository")
		}
	}

	return h.HandleError(fmt.Errorf("repository %s not found", repositoryName), "get_repository")
}

func (h *Handler) ListRepositoryTags(registryName, repositoryName string) (*mcp_golang.ToolResponse, error) {
	client := h.doClient.GetClient()
	
	tags, _, err := client.Registry.ListRepositoryTags(context.Background(), registryName, repositoryName, &godo.ListOptions{})
	if err != nil {
		return h.HandleError(err, "list_repository_tags")
	}

	return h.HandleSuccess(tags, "list_repository_tags")
}