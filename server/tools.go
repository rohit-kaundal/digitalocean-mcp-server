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
		// Test connection
		{
			Name:        "test_connection",
			Description: "Test connection to DigitalOcean API",
			Handler: func(arguments types.EmptyArgs) (*mcp_golang.ToolResponse, error) {
				return handler.TestConnection()
			},
		},
		
		// Droplet tools
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
			Name:        "resize_droplet",
			Description: "Resize a droplet to a different size",
			Handler: func(arguments types.ResizeDropletArgs) (*mcp_golang.ToolResponse, error) {
				return handler.ResizeDroplet(arguments.DropletID, arguments.Size, arguments.Disk)
			},
		},
		{
			Name:        "create_droplet_snapshot",
			Description: "Create a snapshot of a droplet",
			Handler: func(arguments types.CreateDropletSnapshotArgs) (*mcp_golang.ToolResponse, error) {
				return handler.CreateDropletSnapshot(arguments.DropletID, arguments.Name)
			},
		},
		
		// Volume tools
		{
			Name:        "list_volumes",
			Description: "List all volumes in the account",
			Handler: func(arguments types.ListVolumesArgs) (*mcp_golang.ToolResponse, error) {
				return handler.ListVolumes(arguments.Region)
			},
		},
		{
			Name:        "get_volume",
			Description: "Get details of a specific volume",
			Handler: func(arguments types.GetVolumeArgs) (*mcp_golang.ToolResponse, error) {
				return handler.GetVolume(arguments.VolumeID)
			},
		},
		{
			Name:        "create_volume",
			Description: "Create a new volume",
			Handler: func(arguments types.CreateVolumeArgs) (*mcp_golang.ToolResponse, error) {
				return handler.CreateVolume(arguments.Name, arguments.Region, arguments.SizeGigaBytes, arguments.Description)
			},
		},
		{
			Name:        "delete_volume",
			Description: "Delete a volume",
			Handler: func(arguments types.DeleteVolumeArgs) (*mcp_golang.ToolResponse, error) {
				return handler.DeleteVolume(arguments.VolumeID)
			},
		},
		{
			Name:        "attach_volume",
			Description: "Attach a volume to a droplet",
			Handler: func(arguments types.AttachVolumeArgs) (*mcp_golang.ToolResponse, error) {
				return handler.AttachVolume(arguments.VolumeID, arguments.DropletID)
			},
		},
		{
			Name:        "detach_volume",
			Description: "Detach a volume from a droplet",
			Handler: func(arguments types.DetachVolumeArgs) (*mcp_golang.ToolResponse, error) {
				return handler.DetachVolume(arguments.VolumeID, arguments.DropletID)
			},
		},
		{
			Name:        "resize_volume",
			Description: "Resize a volume",
			Handler: func(arguments types.ResizeVolumeArgs) (*mcp_golang.ToolResponse, error) {
				return handler.ResizeVolume(arguments.VolumeID, arguments.SizeGigaBytes, arguments.Region)
			},
		},
		{
			Name:        "create_volume_snapshot",
			Description: "Create a snapshot of a volume",
			Handler: func(arguments types.CreateVolumeSnapshotArgs) (*mcp_golang.ToolResponse, error) {
				return handler.CreateVolumeSnapshot(arguments.VolumeID, arguments.Name, arguments.Description)
			},
		},
		
		// Snapshot tools
		{
			Name:        "list_snapshots",
			Description: "List all snapshots",
			Handler: func(arguments types.ListSnapshotsArgs) (*mcp_golang.ToolResponse, error) {
				return handler.ListSnapshots(arguments.ResourceType)
			},
		},
		{
			Name:        "list_volume_snapshots",
			Description: "List all volume snapshots",
			Handler: func(arguments types.EmptyArgs) (*mcp_golang.ToolResponse, error) {
				return handler.ListVolumeSnapshots()
			},
		},
		{
			Name:        "list_droplet_snapshots",
			Description: "List all droplet snapshots",
			Handler: func(arguments types.EmptyArgs) (*mcp_golang.ToolResponse, error) {
				return handler.ListDropletSnapshots()
			},
		},
		{
			Name:        "get_snapshot",
			Description: "Get details of a specific snapshot",
			Handler: func(arguments types.GetSnapshotArgs) (*mcp_golang.ToolResponse, error) {
				return handler.GetSnapshot(arguments.SnapshotID)
			},
		},
		{
			Name:        "delete_snapshot",
			Description: "Delete a snapshot",
			Handler: func(arguments types.DeleteSnapshotArgs) (*mcp_golang.ToolResponse, error) {
				return handler.DeleteSnapshot(arguments.SnapshotID)
			},
		},
		
		// Image tools
		{
			Name:        "list_images",
			Description: "List all images",
			Handler: func(arguments types.ListImagesArgs) (*mcp_golang.ToolResponse, error) {
				return handler.ListImages(arguments.Type, arguments.IsPublic)
			},
		},
		{
			Name:        "get_image",
			Description: "Get details of a specific image",
			Handler: func(arguments types.GetImageArgs) (*mcp_golang.ToolResponse, error) {
				return handler.GetImage(arguments.ImageID)
			},
		},
		{
			Name:        "update_image",
			Description: "Update an image",
			Handler: func(arguments types.UpdateImageArgs) (*mcp_golang.ToolResponse, error) {
				return handler.UpdateImage(arguments.ImageID, arguments.Name)
			},
		},
		{
			Name:        "delete_image",
			Description: "Delete an image",
			Handler: func(arguments types.DeleteImageArgs) (*mcp_golang.ToolResponse, error) {
				return handler.DeleteImage(arguments.ImageID)
			},
		},
		{
			Name:        "transfer_image",
			Description: "Transfer an image to another region",
			Handler: func(arguments types.TransferImageArgs) (*mcp_golang.ToolResponse, error) {
				return handler.TransferImage(arguments.ImageID, arguments.RegionSlug)
			},
		},
		{
			Name:        "convert_image_to_snapshot",
			Description: "Convert an image to snapshot",
			Handler: func(arguments types.ConvertImageToSnapshotArgs) (*mcp_golang.ToolResponse, error) {
				return handler.ConvertImageToSnapshot(arguments.ImageID)
			},
		},
		
		// Floating IP tools
		{
			Name:        "list_floating_ips",
			Description: "List all floating IPs",
			Handler: func(arguments types.EmptyArgs) (*mcp_golang.ToolResponse, error) {
				return handler.ListFloatingIPs()
			},
		},
		{
			Name:        "get_floating_ip",
			Description: "Get details of a specific floating IP",
			Handler: func(arguments types.GetFloatingIPArgs) (*mcp_golang.ToolResponse, error) {
				return handler.GetFloatingIP(arguments.IP)
			},
		},
		{
			Name:        "create_floating_ip",
			Description: "Create a new floating IP",
			Handler: func(arguments types.CreateFloatingIPArgs) (*mcp_golang.ToolResponse, error) {
				return handler.CreateFloatingIP(arguments.Region, arguments.DropletID)
			},
		},
		{
			Name:        "delete_floating_ip",
			Description: "Delete a floating IP",
			Handler: func(arguments types.DeleteFloatingIPArgs) (*mcp_golang.ToolResponse, error) {
				return handler.DeleteFloatingIP(arguments.IP)
			},
		},
		{
			Name:        "assign_floating_ip",
			Description: "Assign a floating IP to a droplet",
			Handler: func(arguments types.AssignFloatingIPArgs) (*mcp_golang.ToolResponse, error) {
				return handler.AssignFloatingIP(arguments.IP, arguments.DropletID)
			},
		},
		{
			Name:        "unassign_floating_ip",
			Description: "Unassign a floating IP from a droplet",
			Handler: func(arguments types.UnassignFloatingIPArgs) (*mcp_golang.ToolResponse, error) {
				return handler.UnassignFloatingIP(arguments.IP)
			},
		},
		
		// Load Balancer tools
		{
			Name:        "list_load_balancers",
			Description: "List all load balancers",
			Handler: func(arguments types.EmptyArgs) (*mcp_golang.ToolResponse, error) {
				return handler.ListLoadBalancers()
			},
		},
		{
			Name:        "get_load_balancer",
			Description: "Get details of a specific load balancer",
			Handler: func(arguments types.GetLoadBalancerArgs) (*mcp_golang.ToolResponse, error) {
				return handler.GetLoadBalancer(arguments.LoadBalancerID)
			},
		},
		{
			Name:        "create_load_balancer",
			Description: "Create a new load balancer",
			Handler: func(arguments types.CreateLoadBalancerArgs) (*mcp_golang.ToolResponse, error) {
				return handler.CreateLoadBalancer(arguments.Name, arguments.Algorithm, arguments.Region, arguments.ForwardingRules, arguments.DropletIDs)
			},
		},
		{
			Name:        "update_load_balancer",
			Description: "Update a load balancer",
			Handler: func(arguments types.UpdateLoadBalancerArgs) (*mcp_golang.ToolResponse, error) {
				return handler.UpdateLoadBalancer(arguments.LoadBalancerID, arguments.Name, arguments.Algorithm, arguments.Region, arguments.ForwardingRules, arguments.DropletIDs)
			},
		},
		{
			Name:        "delete_load_balancer",
			Description: "Delete a load balancer",
			Handler: func(arguments types.DeleteLoadBalancerArgs) (*mcp_golang.ToolResponse, error) {
				return handler.DeleteLoadBalancer(arguments.LoadBalancerID)
			},
		},
		{
			Name:        "add_droplets_to_load_balancer",
			Description: "Add droplets to a load balancer",
			Handler: func(arguments types.AddDropletsToLoadBalancerArgs) (*mcp_golang.ToolResponse, error) {
				return handler.AddDropletsToLoadBalancer(arguments.LoadBalancerID, arguments.DropletIDs)
			},
		},
		{
			Name:        "remove_droplets_from_load_balancer",
			Description: "Remove droplets from a load balancer",
			Handler: func(arguments types.RemoveDropletsFromLoadBalancerArgs) (*mcp_golang.ToolResponse, error) {
				return handler.RemoveDropletsFromLoadBalancer(arguments.LoadBalancerID, arguments.DropletIDs)
			},
		},
		{
			Name:        "add_forwarding_rules_to_load_balancer",
			Description: "Add forwarding rules to a load balancer",
			Handler: func(arguments types.AddForwardingRulesToLoadBalancerArgs) (*mcp_golang.ToolResponse, error) {
				return handler.AddForwardingRulesToLoadBalancer(arguments.LoadBalancerID, arguments.ForwardingRules)
			},
		},
		{
			Name:        "remove_forwarding_rules_from_load_balancer",
			Description: "Remove forwarding rules from a load balancer",
			Handler: func(arguments types.RemoveForwardingRulesFromLoadBalancerArgs) (*mcp_golang.ToolResponse, error) {
				return handler.RemoveForwardingRulesFromLoadBalancer(arguments.LoadBalancerID, arguments.ForwardingRules)
			},
		},
		
		// Firewall tools
		{
			Name:        "list_firewalls",
			Description: "List all firewalls",
			Handler: func(arguments types.EmptyArgs) (*mcp_golang.ToolResponse, error) {
				return handler.ListFirewalls()
			},
		},
		{
			Name:        "get_firewall",
			Description: "Get details of a specific firewall",
			Handler: func(arguments types.GetFirewallArgs) (*mcp_golang.ToolResponse, error) {
				return handler.GetFirewall(arguments.FirewallID)
			},
		},
		{
			Name:        "create_firewall",
			Description: "Create a new firewall",
			Handler: func(arguments types.CreateFirewallArgs) (*mcp_golang.ToolResponse, error) {
				return handler.CreateFirewall(arguments.Name, arguments.InboundRules, arguments.OutboundRules, arguments.DropletIDs, arguments.Tags)
			},
		},
		{
			Name:        "update_firewall",
			Description: "Update a firewall",
			Handler: func(arguments types.UpdateFirewallArgs) (*mcp_golang.ToolResponse, error) {
				return handler.UpdateFirewall(arguments.FirewallID, arguments.Name, arguments.InboundRules, arguments.OutboundRules)
			},
		},
		{
			Name:        "delete_firewall",
			Description: "Delete a firewall",
			Handler: func(arguments types.DeleteFirewallArgs) (*mcp_golang.ToolResponse, error) {
				return handler.DeleteFirewall(arguments.FirewallID)
			},
		},
		{
			Name:        "add_droplets_to_firewall",
			Description: "Add droplets to a firewall",
			Handler: func(arguments types.AddDropletsToFirewallArgs) (*mcp_golang.ToolResponse, error) {
				return handler.AddDropletsToFirewall(arguments.FirewallID, arguments.DropletIDs)
			},
		},
		{
			Name:        "remove_droplets_from_firewall",
			Description: "Remove droplets from a firewall",
			Handler: func(arguments types.RemoveDropletsFromFirewallArgs) (*mcp_golang.ToolResponse, error) {
				return handler.RemoveDropletsFromFirewall(arguments.FirewallID, arguments.DropletIDs)
			},
		},
		{
			Name:        "add_tags_to_firewall",
			Description: "Add tags to a firewall",
			Handler: func(arguments types.AddTagsToFirewallArgs) (*mcp_golang.ToolResponse, error) {
				return handler.AddTagsToFirewall(arguments.FirewallID, arguments.Tags)
			},
		},
		{
			Name:        "remove_tags_from_firewall",
			Description: "Remove tags from a firewall",
			Handler: func(arguments types.RemoveTagsFromFirewallArgs) (*mcp_golang.ToolResponse, error) {
				return handler.RemoveTagsFromFirewall(arguments.FirewallID, arguments.Tags)
			},
		},
		{
			Name:        "add_rules_to_firewall",
			Description: "Add rules to a firewall",
			Handler: func(arguments types.AddRulesToFirewallArgs) (*mcp_golang.ToolResponse, error) {
				return handler.AddRulesToFirewall(arguments.FirewallID, arguments.InboundRules, arguments.OutboundRules)
			},
		},
		{
			Name:        "remove_rules_from_firewall",
			Description: "Remove rules from a firewall",
			Handler: func(arguments types.RemoveRulesFromFirewallArgs) (*mcp_golang.ToolResponse, error) {
				return handler.RemoveRulesFromFirewall(arguments.FirewallID, arguments.InboundRules, arguments.OutboundRules)
			},
		},
		
		// Registry tools
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
		
		// Kubernetes tools
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