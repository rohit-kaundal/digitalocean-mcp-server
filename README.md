# DigitalOcean MCP Server

A comprehensive Model Context Protocol (MCP) server that provides programmatic access to DigitalOcean's API. This server exposes **48 tools** across **7 major service categories** for complete infrastructure management through the MCP interface.

## Features

- **🖥️ Droplet Management**: Complete lifecycle management with resize and snapshot capabilities
- **💾 Volume Management**: Block storage operations including attach/detach and snapshots
- **📸 Snapshot Operations**: Backup and restore functionality for droplets and volumes
- **🖼️ Image Management**: Custom image operations, transfers, and conversions
- **🌐 Floating IP Management**: Static IP allocation, assignment, and management
- **⚖️ Load Balancer Operations**: Traffic distribution with full CRUD operations
- **🔥 Firewall Management**: Complete network security with rule and policy management
- **☸️ Kubernetes Operations**: Comprehensive cluster and node pool management
- **📦 Container Registry**: Access and manage DigitalOcean container registries
- **✅ Connection Testing**: Verify API connectivity and authentication

## Prerequisites

- Go 1.19 or higher
- **DigitalOcean API token** (required - see Configuration section)

## Installation

### From Source

```bash
git clone https://github.com/rohit-kaundal/digitalocean-mcp-server.git
cd digitalocean-mcp-server
go mod tidy
go build -o digitalocean-mcp-server
```

### Using Go Install

```bash
go install github.com/rohit-kaundal/digitalocean-mcp-server@latest
```

## Configuration

### Environment Variables

**REQUIRED**: Set your DigitalOcean API token before running the server:

```bash
export DIGITALOCEAN_ACCESS_TOKEN="your_digitalocean_api_token"
```

⚠️ **The server will not function without this environment variable set.**

### Getting a DigitalOcean API Token

1. Log in to your [DigitalOcean Control Panel](https://cloud.digitalocean.com/)
2. Navigate to **API** in the left sidebar
3. Click **Generate New Token**
4. Give your token a name and select appropriate scopes
5. Copy the generated token (you won't be able to see it again)

## Usage

### Running the Server

First, ensure your DigitalOcean API token is set:

```bash
export DIGITALOCEAN_ACCESS_TOKEN="your_digitalocean_api_token"
```

Then run the server:

```bash
./digitalocean-mcp-server
```

Or run directly with Go:

```bash
go run main.go
```

The server will start and listen for MCP requests via stdio transport.

### Available Tools (48 Total)

#### Connection & Testing
- **`test_connection`** - Test API connectivity and authentication

#### Droplet Management (7 tools)
- **`list_droplets`** - List all droplets with pagination support
- **`get_droplet`** - Get detailed information about a specific droplet
- **`create_droplet`** - Create a new droplet with custom specifications
- **`delete_droplet`** - Permanently delete a droplet
- **`resize_droplet`** - Resize droplet to different size (CPU/RAM/disk)
- **`create_droplet_snapshot`** - Create a snapshot backup of a droplet

#### Volume Management (8 tools)
- **`list_volumes`** - List all block storage volumes (optionally by region)
- **`get_volume`** - Get detailed volume information
- **`create_volume`** - Create new block storage volume
- **`delete_volume`** - Delete a volume
- **`attach_volume`** - Attach volume to a droplet
- **`detach_volume`** - Detach volume from a droplet
- **`resize_volume`** - Expand volume storage capacity
- **`create_volume_snapshot`** - Create snapshot backup of a volume

#### Snapshot Operations (6 tools)
- **`list_snapshots`** - List all snapshots (filter by droplet/volume)
- **`list_volume_snapshots`** - List volume-specific snapshots
- **`list_droplet_snapshots`** - List droplet-specific snapshots
- **`get_snapshot`** - Get detailed snapshot information
- **`delete_snapshot`** - Delete a snapshot

#### Image Management (6 tools)
- **`list_images`** - List available images (distribution/application/user)
- **`get_image`** - Get image details by ID or slug
- **`update_image`** - Update image metadata (name, description)
- **`delete_image`** - Delete custom images
- **`transfer_image`** - Transfer image to different region
- **`convert_image_to_snapshot`** - Convert image to snapshot format

#### Floating IP Management (6 tools)
- **`list_floating_ips`** - List all floating IP addresses
- **`get_floating_ip`** - Get floating IP details and assignment status
- **`create_floating_ip`** - Create new floating IP (regional or assigned)
- **`delete_floating_ip`** - Release floating IP
- **`assign_floating_ip`** - Assign floating IP to droplet
- **`unassign_floating_ip`** - Unassign floating IP from droplet

#### Load Balancer Operations (9 tools)
- **`list_load_balancers`** - List all load balancers
- **`get_load_balancer`** - Get load balancer configuration and status
- **`create_load_balancer`** - Create new load balancer with forwarding rules
- **`update_load_balancer`** - Update load balancer configuration
- **`delete_load_balancer`** - Delete load balancer
- **`add_droplets_to_load_balancer`** - Add droplets to load balancer pool
- **`remove_droplets_from_load_balancer`** - Remove droplets from pool
- **`add_forwarding_rules_to_load_balancer`** - Add traffic forwarding rules
- **`remove_forwarding_rules_from_load_balancer`** - Remove forwarding rules

#### Firewall Management (11 tools)
- **`list_firewalls`** - List all firewalls in the account
- **`get_firewall`** - Get detailed firewall configuration and rules
- **`create_firewall`** - Create new firewall with inbound/outbound rules
- **`update_firewall`** - Update firewall configuration and rules
- **`delete_firewall`** - Remove firewall from account
- **`add_droplets_to_firewall`** - Assign droplets to firewall protection
- **`remove_droplets_from_firewall`** - Remove droplets from firewall
- **`add_tags_to_firewall`** - Add tags to firewall for organization
- **`remove_tags_from_firewall`** - Remove tags from firewall
- **`add_rules_to_firewall`** - Add new security rules to firewall
- **`remove_rules_from_firewall`** - Remove existing security rules

#### Kubernetes Clusters (4 tools)
- **`list_k8s_clusters`** - List all Kubernetes clusters
- **`get_k8s_cluster`** - Get cluster details and status
- **`create_k8s_cluster`** - Create new Kubernetes cluster
- **`delete_k8s_cluster`** - Delete Kubernetes cluster

#### Container Registry (2 tools)
- **`list_registries`** - List all container registries
- **`get_registry`** - Get registry details and repositories

### Example MCP Client Usage

#### Basic Operations
```json
{
  "method": "tools/call",
  "params": {
    "name": "list_droplets",
    "arguments": {
      "page": 1,
      "per_page": 25
    }
  }
}
```

```json
{
  "method": "tools/call",
  "params": {
    "name": "create_droplet",
    "arguments": {
      "name": "my-server",
      "region": "nyc3",
      "size": "s-1vcpu-1gb",
      "image": "ubuntu-22-04-x64"
    }
  }
}
```

#### Volume Management
```json
{
  "method": "tools/call",
  "params": {
    "name": "create_volume",
    "arguments": {
      "name": "my-storage",
      "region": "nyc3",
      "size_gigabytes": 100,
      "description": "Additional storage for applications"
    }
  }
}
```

```json
{
  "method": "tools/call",
  "params": {
    "name": "attach_volume",
    "arguments": {
      "volume_id": "volume-123",
      "droplet_id": 456
    }
  }
}
```

#### Load Balancer Setup
```json
{
  "method": "tools/call",
  "params": {
    "name": "create_load_balancer",
    "arguments": {
      "name": "web-lb",
      "algorithm": "round_robin",
      "region": "nyc3",
      "forwarding_rules": [
        {
          "entry_protocol": "http",
          "entry_port": 80,
          "target_protocol": "http",
          "target_port": 8080
        }
      ],
      "droplet_ids": [123, 456]
    }
  }
}
```

#### Floating IP Management
```json
{
  "method": "tools/call",
  "params": {
    "name": "create_floating_ip",
    "arguments": {
      "region": "nyc3"
    }
  }
}
```

```json
{
  "method": "tools/call",
  "params": {
    "name": "assign_floating_ip",
    "arguments": {
      "ip": "192.168.1.100",
      "droplet_id": 123
    }
  }
}
```

#### Firewall Configuration
```json
{
  "method": "tools/call",
  "params": {
    "name": "create_firewall",
    "arguments": {
      "name": "web-firewall",
      "inbound_rules": [
        {
          "protocol": "tcp",
          "ports": "80",
          "sources": {
            "addresses": ["0.0.0.0/0"]
          }
        },
        {
          "protocol": "tcp",
          "ports": "443",
          "sources": {
            "addresses": ["0.0.0.0/0"]
          }
        }
      ],
      "outbound_rules": [
        {
          "protocol": "tcp",
          "ports": "all",
          "destinations": {
            "addresses": ["0.0.0.0/0"]
          }
        }
      ],
      "droplet_ids": [123, 456]
    }
  }
}
```

```json
{
  "method": "tools/call",
  "params": {
    "name": "add_droplets_to_firewall",
    "arguments": {
      "firewall_id": "firewall-123",
      "droplet_ids": [789, 101112]
    }
  }
}
```

## Development

### Project Structure

```
digitalocean-mcp-server/
├── main.go                 # Entry point
├── server/
│   ├── server.go          # MCP server initialization
│   └── tools.go           # Tool registration
├── client/
│   └── digitalocean.go    # DigitalOcean API client
├── handlers/
│   ├── common.go          # Shared handler functionality
│   ├── droplets.go        # Droplet operations
│   ├── volumes.go         # Volume operations
│   ├── snapshots.go       # Snapshot operations
│   ├── images.go          # Image operations
│   ├── floating_ips.go    # Floating IP operations
│   ├── load_balancers.go  # Load balancer operations
│   ├── firewalls.go       # Firewall operations
│   ├── kubernetes.go      # Kubernetes operations
│   └── registry.go        # Registry operations
├── types/
│   └── args.go            # Request argument types
└── CLAUDE.md              # AI assistant instructions
```

### Development Commands

```bash
# Run the server
go run main.go

# Run tests
go test ./...

# Build binary
go build -o digitalocean-mcp-server

# Update dependencies
go mod tidy
```

### Adding New Tools

1. Define argument types in `types/args.go`
2. Implement handler logic in the appropriate file under `handlers/`
3. Register the tool in `server/tools.go`
4. Update this documentation

## Dependencies

- [godo](https://github.com/digitalocean/godo) - DigitalOcean API client
- [mcp-golang](https://github.com/metoro-io/mcp-golang) - MCP protocol implementation
- [oauth2](https://golang.org/x/oauth2) - OAuth2 authentication

## Error Handling

All tools return standardized MCP responses:

- **Success**: JSON-formatted data with the requested information
- **Error**: Structured error messages with details about what went wrong

Common error scenarios:
- Invalid or missing API token
- Insufficient permissions
- Resource not found
- API rate limiting
- Network connectivity issues

## Contributing

We welcome contributions! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Support

- Report issues on [GitHub Issues](https://github.com/rohit-kaundal/digitalocean-mcp-server/issues)
- Check [DigitalOcean API Documentation](https://docs.digitalocean.com/reference/api/)
- Review [MCP Protocol Specification](https://modelcontextprotocol.io/docs)

## Changelog

### v2.0.1
- **Firewall Management**: Added comprehensive network security with 11 firewall tools
- **Security Rules**: Complete inbound/outbound rule management and configuration
- **Droplet Protection**: Assign and manage firewall protection for droplets
- **Tag Management**: Organize firewalls with tag-based categorization
- **Rule Modification**: Dynamic addition and removal of security rules
- **Enhanced Documentation**: Updated with firewall management examples

### v2.0.0
- **MAJOR EXPANSION**: Added 30+ new API tools across 5 additional service categories
- **Volume Management**: Complete block storage lifecycle (8 tools)
- **Snapshot Operations**: Backup and restore functionality (6 tools)
- **Image Management**: Custom image operations and transfers (6 tools)
- **Floating IP Management**: Static IP allocation and assignment (6 tools)
- **Load Balancer Operations**: Traffic distribution with full CRUD (9 tools)
- **Enhanced Droplets**: Added resize and snapshot capabilities
- **Comprehensive Documentation**: Updated with detailed API coverage
- **Improved Architecture**: Modular handler design for scalability

### v1.0.0
- Initial release
- Droplet management tools (5 tools)
- Kubernetes cluster tools (4 tools)
- Container registry tools (2 tools)
- Connection testing