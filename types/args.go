package types

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