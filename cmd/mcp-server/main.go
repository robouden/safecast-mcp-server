package main

import (
	"context"
	"log"
	"net/http"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func main() {
	// Create MCP server
	mcpServer := server.NewMCPServer(
		"safecast-mcp",
		"0.1.0",
	)

	// Register first MCP tool
	mcpServer.AddTool(
		mcp.Tool{
			Name:        "ping",
			Description: "Health check tool",
		},
		func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			return &mcp.CallToolResult{
				Content: []mcp.Content{
					mcp.TextContent{
						Text: "pong",
					},
				},
			}, nil
		},
	)

	// Wrap in SSE server
	sseServer := server.NewSSEServer(mcpServer)

	// Expose over HTTP
	http.Handle("/sse/", http.StripPrefix("/sse", sseServer))

	log.Println("Starting MCP SSE server on :3333")
	if err := http.ListenAndServe(":3333", nil); err != nil {
		log.Fatal(err)
	}
}
