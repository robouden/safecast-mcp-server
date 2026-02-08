# Conversation Notes: Fix MCP SSE Server 404 on POST

**Date:** 2026-02-08

## Problem

- Opening the SSE stream at `/sse` worked and returned a session ID
- POSTing MCP messages (`initialize`, `tools/list`, `tools/call`) returned **404 page not found**

## Root Cause

In `cmd/mcp-server/main.go`, the SSE server was mounted incorrectly:

```go
http.Handle("/sse/", http.StripPrefix("/sse", sseServer))
```

The `mcp-go` library's `SSEServer.ServeHTTP` routes internally using exact path matching:
- `GET /sse` -> SSE stream handler
- `POST /message` -> message handler

With `StripPrefix("/sse", ...)`, when a client hit `GET /sse/sse`, the prefix was stripped and the handler saw `/sse` — so the SSE stream worked. But the message endpoint URL sent back to the client was `/message?sessionId=...` (no `/sse` prefix), which didn't match any route on the default mux, resulting in a 404.

## Fix Applied

Replaced the manual mux registration with `sseServer.Start()`, which sets the SSEServer as the root HTTP handler for both endpoints:

```go
// Before (broken)
sseServer := server.NewSSEServer(mcpServer)
http.Handle("/sse/", http.StripPrefix("/sse", sseServer))
http.ListenAndServe(":3333", nil)

// After (fixed)
sseServer := server.NewSSEServer(mcpServer,
    server.WithBaseURL("http://localhost:3333"),
)
sseServer.Start(":3333")
```

This lets the SSEServer handle routing itself:
- `GET /sse` -> SSE stream (returns session ID + message endpoint URL)
- `POST /message?sessionId=...` -> MCP message handler

The unused `net/http` import was also removed.

## README Update

Replaced the "What's not working" section with:
- Endpoint reference table (`GET /sse`, `POST /message`)
- Step-by-step curl testing instructions for the full MCP flow (connect, initialize, tools/list, tools/call)

## Files Changed

- `cmd/mcp-server/main.go` — fixed server startup and routing
- `README.md` — updated documentation with working endpoints and test instructions
