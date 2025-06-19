package client

import (
	"github.com/KunalSin9h/modern-mcp-go/client/transport"
	"github.com/KunalSin9h/modern-mcp-go/server"
)

// NewInProcessClient connect directly to a mcp server object in the same process
func NewInProcessClient(server *server.MCPServer) (*Client, error) {
	inProcessTransport := transport.NewInProcessTransport(server)
	return NewClient(inProcessTransport), nil
}
