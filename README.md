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
```

You should see:
Starting MCP SSE server on :3333

## Endpoints

| Method | Path | Description |
|--------|------|-------------|
| GET | `/sse` | Opens an SSE stream, returns a session ID and message endpoint URL |
| POST | `/message?sessionId=...` | Send MCP JSON-RPC messages (`initialize`, `tools/list`, `tools/call`) |

## Testing with curl

1. Open the SSE stream in one terminal:

```bash
curl -N http://localhost:3333/sse
```

You'll receive an `endpoint` event with the message URL.

2. In another terminal, POST an MCP `initialize` request (replace the session ID):

```bash
curl -X POST "http://localhost:3333/message?sessionId=<SESSION_ID>" \
  -H "Content-Type: application/json" \
  -d '{"jsonrpc":"2.0","id":1,"method":"initialize","params":{"protocolVersion":"2024-11-05","capabilities":{},"clientInfo":{"name":"test","version":"0.1.0"}}}'
```

3. List available tools:

```bash
curl -X POST "http://localhost:3333/message?sessionId=<SESSION_ID>" \
  -H "Content-Type: application/json" \
  -d '{"jsonrpc":"2.0","id":2,"method":"tools/list"}'
```

4. Call the `ping` tool:

```bash
curl -X POST "http://localhost:3333/message?sessionId=<SESSION_ID>" \
  -H "Content-Type: application/json" \
  -d '{"jsonrpc":"2.0","id":3,"method":"tools/call","params":{"name":"ping"}}'
```

Responses are delivered via the SSE stream in terminal 1.
