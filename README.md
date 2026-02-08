# Safecast MCP Server (Go)

Minimal Go-based MCP server using `mcp-go` with SSE transport.  
Currently includes a simple `ping → pong` tool to validate MCP wiring.

## Prerequisites

- Go 1.21+
- macOS / Linux
- No other dependencies

---

## Run the server

```bash
git clone https://github.com/van-van-nguyen/safecast-mcp-server.git
cd safecast-mcp-server
go run cmd/mcp-server/main.go

You should see:
Starting MCP SSE server on :3333

## What works

- Server boots cleanly
- MCP server initializes
- `ping` tool is registered in code
- SSE endpoint responds

---

## What’s not working (current issue)

When testing manually via `curl`:

- Opening the SSE stream works and returns a session ID
- POSTing MCP messages (`initialize`, `tools/list`, `tools/call`) returns `404 page not found`
