package types

import "github.com/digitalocean/godo"

type EmptyArgs struct{}

type ListDropletsArgs struct {
	Page    int `json:"page" jsonschema:"description=Page number to retrieve (starting from 1),default=1"`
	PerPage int `json:"per_page" jsonschema:"description=Number of items per page (1-200),default=25"`
}

type GetDropletArgs struct {
	DropletID int `json:"droplet_id" jsonschema:"description=ID of the droplet to retrieve"`
}

type CreateDropletArgs struct {
	Name   string `json:"name" jsonschema:"description=Name of the droplet"`
	Region string `json:"region" jsonschema:"description=Region slug (e.g., 'nyc3', 'sfo2')"`
	Size   string `json:"size" jsonschema:"description=Size slug (e.g., 's-1vcpu-1gb')"`
	Image  string `json:"image" jsonschema:"description=Image slug (e.g., 'ubuntu-22-04-x64')"`
}

type DeleteDropletArgs struct {
	DropletID int `json:"droplet_id" jsonschema:"description=ID of the droplet to delete"`
}

type ResizeDropletArgs struct {
	DropletID int    `json:"droplet_id" jsonschema:"description=ID of the droplet to resize"`
	Size      string `json:"size" jsonschema:"description=New size slug (e.g., 's-2vcpu-2gb')"`
	Disk      bool   `json:"disk" jsonschema:"description=Whether to resize disk (permanent, cannot be undone),default=false"`
}

type CreateDropletSnapshotArgs struct {
	DropletID int    `json:"droplet_id" jsonschema:"description=ID of the droplet to snapshot"`
	Name      string `json:"name" jsonschema:"description=Name for the snapshot"`
}

type GetRegistryArgs struct {
	RegistryName string `json:"registry_name" jsonschema:"description=Name of the registry"`
}

type GetK8SClusterArgs struct {
	ClusterID string `json:"cluster_id" jsonschema:"description=ID of the cluster"`
}

type CreateK8SClusterArgs struct {
	Name         string `json:"name" jsonschema:"description=Name of the cluster"`
	Region       string `json:"region" jsonschema:"description=Region slug (e.g., 'nyc3', 'sfo2')"`
	Version      string `json:"version" jsonschema:"description=Kubernetes version (e.g., '1.28.2-do.0')"`
	NodePoolSize string `json:"node_pool_size" jsonschema:"description=Node pool size (e.g., 's-2vcpu-2gb')"`
	NodeCount    int    `json:"node_count" jsonschema:"description=Number of nodes in the pool"`
}

type DeleteK8SClusterArgs struct {
	ClusterID string `json:"cluster_id" jsonschema:"description=ID of the cluster to delete"`
}

// Volume-related args
type ListVolumesArgs struct {
	Region string `json:"region,omitempty" jsonschema:"description=Filter volumes by region (optional)"`
}

type GetVolumeArgs struct {
	VolumeID string `json:"volume_id" jsonschema:"description=ID of the volume to retrieve"`
}

type CreateVolumeArgs struct {
	Name          string `json:"name" jsonschema:"description=Name of the volume"`
	Region        string `json:"region" jsonschema:"description=Region slug (e.g., 'nyc3', 'sfo2')"`
	SizeGigaBytes int64  `json:"size_gigabytes" jsonschema:"description=Size of the volume in gigabytes"`
	Description   string `json:"description,omitempty" jsonschema:"description=Description of the volume (optional)"`
}

type DeleteVolumeArgs struct {
	VolumeID string `json:"volume_id" jsonschema:"description=ID of the volume to delete"`
}

type AttachVolumeArgs struct {
	VolumeID  string `json:"volume_id" jsonschema:"description=ID of the volume to attach"`
	DropletID int    `json:"droplet_id" jsonschema:"description=ID of the droplet to attach to"`
}

type DetachVolumeArgs struct {
	VolumeID  string `json:"volume_id" jsonschema:"description=ID of the volume to detach"`
	DropletID int    `json:"droplet_id" jsonschema:"description=ID of the droplet to detach from"`
}

type ResizeVolumeArgs struct {
	VolumeID      string `json:"volume_id" jsonschema:"description=ID of the volume to resize"`
	SizeGigaBytes int64  `json:"size_gigabytes" jsonschema:"description=New size in gigabytes"`
	Region        string `json:"region" jsonschema:"description=Region slug"`
}

type CreateVolumeSnapshotArgs struct {
	VolumeID    string `json:"volume_id" jsonschema:"description=ID of the volume to snapshot"`
	Name        string `json:"name" jsonschema:"description=Name for the snapshot"`
	Description string `json:"description,omitempty" jsonschema:"description=Description of the snapshot (optional)"`
}

// Snapshot-related args
type ListSnapshotsArgs struct {
	ResourceType string `json:"resource_type,omitempty" jsonschema:"description=Filter by resource type: 'droplet' or 'volume' (optional)"`
}

type GetSnapshotArgs struct {
	SnapshotID string `json:"snapshot_id" jsonschema:"description=ID of the snapshot to retrieve"`
}

type DeleteSnapshotArgs struct {
	SnapshotID string `json:"snapshot_id" jsonschema:"description=ID of the snapshot to delete"`
}

// Image-related args
type ListImagesArgs struct {
	Type     string `json:"type,omitempty" jsonschema:"description=Image type: 'distribution', 'application', 'user' (optional)"`
	IsPublic bool   `json:"is_public,omitempty" jsonschema:"description=Whether to include public images (optional)"`
}

type GetImageArgs struct {
	ImageID string `json:"image_id" jsonschema:"description=ID or slug of the image to retrieve"`
}

type UpdateImageArgs struct {
	ImageID string `json:"image_id" jsonschema:"description=ID of the image to update"`
	Name    string `json:"name" jsonschema:"description=New name for the image"`
}

type DeleteImageArgs struct {
	ImageID string `json:"image_id" jsonschema:"description=ID of the image to delete"`
}

type TransferImageArgs struct {
	ImageID    string `json:"image_id" jsonschema:"description=ID of the image to transfer"`
	RegionSlug string `json:"region_slug" jsonschema:"description=Region slug to transfer to"`
}

type ConvertImageToSnapshotArgs struct {
	ImageID string `json:"image_id" jsonschema:"description=ID of the image to convert"`
}

// Floating IP-related args
type GetFloatingIPArgs struct {
	IP string `json:"ip" jsonschema:"description=Floating IP address"`
}

type CreateFloatingIPArgs struct {
	Region    string `json:"region,omitempty" jsonschema:"description=Region slug for reserved IP (required if no droplet_id)"`
	DropletID int    `json:"droplet_id,omitempty" jsonschema:"description=Droplet ID to assign to (optional)"`
}

type DeleteFloatingIPArgs struct {
	IP string `json:"ip" jsonschema:"description=Floating IP address to delete"`
}

type AssignFloatingIPArgs struct {
	IP        string `json:"ip" jsonschema:"description=Floating IP address"`
	DropletID int    `json:"droplet_id" jsonschema:"description=Droplet ID to assign to"`
}

type UnassignFloatingIPArgs struct {
	IP string `json:"ip" jsonschema:"description=Floating IP address to unassign"`
}

// Load Balancer-related args
type GetLoadBalancerArgs struct {
	LoadBalancerID string `json:"load_balancer_id" jsonschema:"description=ID of the load balancer"`
}

type CreateLoadBalancerArgs struct {
	Name            string                  `json:"name" jsonschema:"description=Name of the load balancer"`
	Algorithm       string                  `json:"algorithm" jsonschema:"description=Load balancing algorithm: 'round_robin', 'least_connections'"`
	Region          string                  `json:"region" jsonschema:"description=Region slug"`
	ForwardingRules []godo.ForwardingRule   `json:"forwarding_rules" jsonschema:"description=Forwarding rules configuration"`
	DropletIDs      []int                   `json:"droplet_ids,omitempty" jsonschema:"description=Droplet IDs to add (optional)"`
}

type UpdateLoadBalancerArgs struct {
	LoadBalancerID  string                  `json:"load_balancer_id" jsonschema:"description=ID of the load balancer"`
	Name            string                  `json:"name" jsonschema:"description=Name of the load balancer"`
	Algorithm       string                  `json:"algorithm" jsonschema:"description=Load balancing algorithm"`
	Region          string                  `json:"region" jsonschema:"description=Region slug"`
	ForwardingRules []godo.ForwardingRule   `json:"forwarding_rules" jsonschema:"description=Forwarding rules configuration"`
	DropletIDs      []int                   `json:"droplet_ids,omitempty" jsonschema:"description=Droplet IDs (optional)"`
}

type DeleteLoadBalancerArgs struct {
	LoadBalancerID string `json:"load_balancer_id" jsonschema:"description=ID of the load balancer to delete"`
}

type AddDropletsToLoadBalancerArgs struct {
	LoadBalancerID string `json:"load_balancer_id" jsonschema:"description=ID of the load balancer"`
	DropletIDs     []int  `json:"droplet_ids" jsonschema:"description=Droplet IDs to add"`
}

type RemoveDropletsFromLoadBalancerArgs struct {
	LoadBalancerID string `json:"load_balancer_id" jsonschema:"description=ID of the load balancer"`
	DropletIDs     []int  `json:"droplet_ids" jsonschema:"description=Droplet IDs to remove"`
}

type AddForwardingRulesToLoadBalancerArgs struct {
	LoadBalancerID  string                `json:"load_balancer_id" jsonschema:"description=ID of the load balancer"`
	ForwardingRules []godo.ForwardingRule `json:"forwarding_rules" jsonschema:"description=Forwarding rules to add"`
}

type RemoveForwardingRulesFromLoadBalancerArgs struct {
	LoadBalancerID  string                `json:"load_balancer_id" jsonschema:"description=ID of the load balancer"`
	ForwardingRules []godo.ForwardingRule `json:"forwarding_rules" jsonschema:"description=Forwarding rules to remove"`
}

// Firewall-related args
type GetFirewallArgs struct {
	FirewallID string `json:"firewall_id" jsonschema:"description=ID of the firewall"`
}

type CreateFirewallArgs struct {
	Name         string                 `json:"name" jsonschema:"description=Name of the firewall"`
	InboundRules []godo.InboundRule     `json:"inbound_rules" jsonschema:"description=Inbound rules configuration"`
	OutboundRules []godo.OutboundRule   `json:"outbound_rules" jsonschema:"description=Outbound rules configuration"`
	DropletIDs   []int                  `json:"droplet_ids,omitempty" jsonschema:"description=Droplet IDs to assign (optional)"`
	Tags         []string               `json:"tags,omitempty" jsonschema:"description=Tags to assign (optional)"`
}

type UpdateFirewallArgs struct {
	FirewallID    string                 `json:"firewall_id" jsonschema:"description=ID of the firewall"`
	Name          string                 `json:"name" jsonschema:"description=Name of the firewall"`
	InboundRules  []godo.InboundRule     `json:"inbound_rules" jsonschema:"description=Inbound rules configuration"`
	OutboundRules []godo.OutboundRule    `json:"outbound_rules" jsonschema:"description=Outbound rules configuration"`
}

type DeleteFirewallArgs struct {
	FirewallID string `json:"firewall_id" jsonschema:"description=ID of the firewall to delete"`
}

type AddDropletsToFirewallArgs struct {
	FirewallID string `json:"firewall_id" jsonschema:"description=ID of the firewall"`
	DropletIDs []int  `json:"droplet_ids" jsonschema:"description=Droplet IDs to add"`
}

type RemoveDropletsFromFirewallArgs struct {
	FirewallID string `json:"firewall_id" jsonschema:"description=ID of the firewall"`
	DropletIDs []int  `json:"droplet_ids" jsonschema:"description=Droplet IDs to remove"`
}

type AddTagsToFirewallArgs struct {
	FirewallID string   `json:"firewall_id" jsonschema:"description=ID of the firewall"`
	Tags       []string `json:"tags" jsonschema:"description=Tags to add"`
}

type RemoveTagsFromFirewallArgs struct {
	FirewallID string   `json:"firewall_id" jsonschema:"description=ID of the firewall"`
	Tags       []string `json:"tags" jsonschema:"description=Tags to remove"`
}

type AddRulesToFirewallArgs struct {
	FirewallID    string              `json:"firewall_id" jsonschema:"description=ID of the firewall"`
	InboundRules  []godo.InboundRule  `json:"inbound_rules,omitempty" jsonschema:"description=Inbound rules to add (optional)"`
	OutboundRules []godo.OutboundRule `json:"outbound_rules,omitempty" jsonschema:"description=Outbound rules to add (optional)"`
}

type RemoveRulesFromFirewallArgs struct {
	FirewallID    string              `json:"firewall_id" jsonschema:"description=ID of the firewall"`
	InboundRules  []godo.InboundRule  `json:"inbound_rules,omitempty" jsonschema:"description=Inbound rules to remove (optional)"`
	OutboundRules []godo.OutboundRule `json:"outbound_rules,omitempty" jsonschema:"description=Outbound rules to remove (optional)"`
}