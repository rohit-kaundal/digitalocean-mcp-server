# DigitalOcean MCP Server

A Model Context Protocol (MCP) server that provides programmatic access to DigitalOcean's API. This server exposes tools for managing droplets, Kubernetes clusters, and container registries through the MCP interface.

## Features

- **Droplet Management**: Create, list, get, and delete DigitalOcean droplets
- **Kubernetes Operations**: Manage DigitalOcean Kubernetes clusters
- **Container Registry**: Access and manage DigitalOcean container registries
- **Connection Testing**: Verify API connectivity and authentication

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

### Available Tools

#### Droplet Management

- **`test_connection`** - Test API connectivity and authentication
- **`list_droplets`** - List all droplets in your account
- **`get_droplet`** - Get details of a specific droplet
  - Parameters: `droplet_id` (integer)
- **`create_droplet`** - Create a new droplet
  - Parameters: `name` (string), `region` (string), `size` (string), `image` (string)
- **`delete_droplet`** - Delete a droplet
  - Parameters: `droplet_id` (integer)

#### Kubernetes Clusters

- **`list_k8s_clusters`** - List all Kubernetes clusters
- **`get_k8s_cluster`** - Get details of a specific cluster
  - Parameters: `cluster_id` (string)
- **`create_k8s_cluster`** - Create a new Kubernetes cluster
  - Parameters: `name` (string), `region` (string), `version` (string), `node_pools` (array)
- **`delete_k8s_cluster`** - Delete a Kubernetes cluster
  - Parameters: `cluster_id` (string)

#### Container Registry

- **`list_registries`** - List all container registries
- **`get_registry`** - Get details of a specific registry
  - Parameters: `registry_name` (string)

### Example MCP Client Usage

```json
{
  "method": "tools/call",
  "params": {
    "name": "list_droplets",
    "arguments": {}
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
      "image": "ubuntu-20-04-x64"
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

### v1.0.0
- Initial release
- Droplet management tools
- Kubernetes cluster tools  
- Container registry tools
- Connection testing