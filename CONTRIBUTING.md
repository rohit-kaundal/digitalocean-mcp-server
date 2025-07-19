# Contributing to DigitalOcean MCP Server

We welcome contributions to the DigitalOcean MCP Server! This document provides guidelines for contributing to the project.

## Table of Contents

- [Code of Conduct](#code-of-conduct)
- [Getting Started](#getting-started)
- [Development Setup](#development-setup)
- [Making Changes](#making-changes)
- [Submitting Changes](#submitting-changes)
- [Coding Standards](#coding-standards)
- [Testing](#testing)
- [Documentation](#documentation)

## Code of Conduct

This project and everyone participating in it is governed by our Code of Conduct. By participating, you are expected to uphold this code. Please be respectful and constructive in all interactions.

## Getting Started

### Prerequisites

- Go 1.19 or higher
- Git
- DigitalOcean API token (for testing)

### Fork and Clone

1. Fork the repository on GitHub
2. Clone your fork locally:

```bash
git clone https://github.com/YOUR_USERNAME/digitalocean-mcp-server.git
cd digitalocean-mcp-server
```

3. Add the upstream repository:

```bash
git remote add upstream https://github.com/rohit-kaundal/digitalocean-mcp-server.git
```

## Development Setup

1. Install dependencies:

```bash
go mod tidy
```

2. Set up your environment:

```bash
export DIGITALOCEAN_ACCESS_TOKEN="your_test_token"
```

3. Verify your setup:

```bash
go run main.go
```

## Making Changes

### Branch Strategy

1. Create a feature branch from `main`:

```bash
git checkout -b feature/your-feature-name
```

2. Make your changes in logical, atomic commits
3. Keep your branch up to date:

```bash
git fetch upstream
git rebase upstream/main
```

### Types of Contributions

#### Bug Fixes

- Include a clear description of the bug
- Add a test case that reproduces the issue
- Fix the issue with minimal changes
- Ensure all tests pass

#### New Features

- Discuss major features in an issue first
- Follow the existing patterns for tool implementation
- Add comprehensive tests
- Update documentation

#### Documentation

- Fix typos, improve clarity
- Add examples and usage scenarios
- Keep documentation in sync with code changes

## Submitting Changes

### Pull Request Process

1. Update documentation as needed
2. Add or update tests for your changes
3. Ensure all tests pass:

```bash
go test ./...
```

4. Run code formatting:

```bash
go fmt ./...
```

5. Create a pull request with:
   - Clear title and description
   - Reference any related issues
   - Include testing instructions

### Pull Request Template

```markdown
## Description
Brief description of changes

## Type of Change
- [ ] Bug fix
- [ ] New feature
- [ ] Documentation update
- [ ] Refactoring

## Testing
- [ ] Tests added/updated
- [ ] All tests pass
- [ ] Manual testing completed

## Checklist
- [ ] Code follows style guidelines
- [ ] Self-review completed
- [ ] Documentation updated
- [ ] No breaking changes (or clearly documented)
```

## Coding Standards

### Go Style Guide

- Follow [Effective Go](https://golang.org/doc/effective_go.html)
- Use `gofmt` for formatting
- Use meaningful variable and function names
- Add comments for exported functions and types

### File Organization

```
handlers/
├── common.go      # Shared functionality
├── droplets.go    # Droplet-specific handlers
├── kubernetes.go  # Kubernetes-specific handlers
└── registry.go    # Registry-specific handlers
```

### Error Handling

- Use the `HandleError` and `HandleSuccess` methods from `handlers/common.go`
- Provide meaningful error messages
- Don't expose internal implementation details in errors

### Tool Implementation Pattern

1. Define arguments in `types/args.go`:

```go
type NewToolArgs struct {
    RequiredParam string `json:"required_param" jsonschema:"required"`
    OptionalParam string `json:"optional_param,omitempty"`
}
```

2. Implement handler:

```go
func HandleNewTool(args NewToolArgs) (*mcp.CallToolResult, error) {
    // Implementation
    return HandleSuccess(result)
}
```

3. Register in `server/tools.go`:

```go
server.RegisterTool("new_tool", "Description", NewToolArgs{}, handlers.HandleNewTool)
```

## Testing

### Unit Tests

- Write tests for all new functionality
- Use table-driven tests where appropriate
- Mock external dependencies

```go
func TestHandleNewTool(t *testing.T) {
    tests := []struct {
        name string
        args NewToolArgs
        want interface{}
        wantErr bool
    }{
        // test cases
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // test implementation
        })
    }
}
```

### Integration Tests

- Test with actual DigitalOcean API (use test resources)
- Clean up resources after tests
- Use environment variables for configuration

### Running Tests

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run specific test
go test -run TestHandleNewTool ./handlers
```

## Documentation

### Code Documentation

- Document all exported functions and types
- Use clear, concise comments
- Include examples where helpful

```go
// HandleNewTool performs a new operation on DigitalOcean resources.
// It takes NewToolArgs and returns the operation result.
//
// Example:
//   result, err := HandleNewTool(NewToolArgs{RequiredParam: "value"})
func HandleNewTool(args NewToolArgs) (*mcp.CallToolResult, error) {
    // implementation
}
```

### README Updates

- Keep tool lists current
- Update examples when APIs change
- Add new configuration options

## Release Process

### Versioning

We use [Semantic Versioning](https://semver.org/):

- **MAJOR**: Breaking changes
- **MINOR**: New features, backward compatible
- **PATCH**: Bug fixes, backward compatible

### Changelog

Update the changelog in README.md with:
- New features
- Bug fixes
- Breaking changes
- Migration instructions

## Getting Help

- Create an issue for questions
- Join discussions in existing issues
- Reference documentation and examples

## Recognition

Contributors will be recognized in:
- GitHub contributors list
- Release notes for significant contributions
- Project documentation

Thank you for contributing to the DigitalOcean MCP Server!