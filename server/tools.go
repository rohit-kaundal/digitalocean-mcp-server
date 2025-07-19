package server

import (
	"digitalocean-mcp-server/handlers"
	"digitalocean-mcp-server/types"
	"log"

	mcp_golang "github.com/metoro-io/mcp-golang"
)

type ToolDefinition struct {
	Name        string
	Description string
	Handler     interface{}
}

func RegisterTools(server *mcp_golang.Server, handler *handlers.Handler) error {
	tools := []ToolDefinition{
		{
			Name:        "test_connection",
			Description: "Test connection to DigitalOcean API",
			Handler: func(arguments types.EmptyArgs) (*mcp_golang.ToolResponse, error) {
				return handler.TestConnection()
			},
		},
		{
			Name:        "list_droplets",
			Description: "List all droplets in the account",
			Handler: func(arguments types.ListDropletsArgs) (*mcp_golang.ToolResponse, error) {
				return handler.ListDroplets(arguments.Page, arguments.PerPage)
			},
		},
		{
			Name:        "get_droplet",
			Description: "Get details of a specific droplet",
			Handler: func(arguments types.GetDropletArgs) (*mcp_golang.ToolResponse, error) {
				return handler.GetDroplet(arguments.DropletID)
			},
		},
		{
			Name:        "create_droplet",
			Description: "Create a new droplet",
			Handler: func(arguments types.CreateDropletArgs) (*mcp_golang.ToolResponse, error) {
				return handler.CreateDroplet(arguments.Name, arguments.Region, arguments.Size, arguments.Image)
			},
		},
		{
			Name:        "delete_droplet",
			Description: "Delete a droplet",
			Handler: func(arguments types.DeleteDropletArgs) (*mcp_golang.ToolResponse, error) {
				return handler.DeleteDroplet(arguments.DropletID)
			},
		},
		{
			Name:        "list_registries",
			Description: "List all container registries",
			Handler: func(arguments types.EmptyArgs) (*mcp_golang.ToolResponse, error) {
				return handler.ListRegistries()
			},
		},
		{
			Name:        "get_registry",
			Description: "Get details of a specific registry",
			Handler: func(arguments types.GetRegistryArgs) (*mcp_golang.ToolResponse, error) {
				return handler.GetRegistry(arguments.RegistryName)
			},
		},
		{
			Name:        "list_k8s_clusters",
			Description: "List all Kubernetes clusters",
			Handler: func(arguments types.EmptyArgs) (*mcp_golang.ToolResponse, error) {
				return handler.ListK8SClusters()
			},
		},
		{
			Name:        "get_k8s_cluster",
			Description: "Get details of a specific Kubernetes cluster",
			Handler: func(arguments types.GetK8SClusterArgs) (*mcp_golang.ToolResponse, error) {
				return handler.GetK8SCluster(arguments.ClusterID)
			},
		},
		{
			Name:        "create_k8s_cluster",
			Description: "Create a new Kubernetes cluster",
			Handler: func(arguments types.CreateK8SClusterArgs) (*mcp_golang.ToolResponse, error) {
				return handler.CreateK8SCluster(arguments.Name, arguments.Region, arguments.Version, arguments.NodePoolSize, arguments.NodeCount)
			},
		},
		{
			Name:        "delete_k8s_cluster",
			Description: "Delete a Kubernetes cluster",
			Handler: func(arguments types.DeleteK8SClusterArgs) (*mcp_golang.ToolResponse, error) {
				return handler.DeleteK8SCluster(arguments.ClusterID)
			},
		},
	}

	for _, tool := range tools {
		if err := server.RegisterTool(tool.Name, tool.Description, tool.Handler); err != nil {
			log.Printf("Failed to register %s tool: %v", tool.Name, err)
			return err
		}
	}

	return nil
}