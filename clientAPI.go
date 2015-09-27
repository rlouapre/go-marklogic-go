package goMarklogicGo

import (
	clients "github.com/ryanjdew/go-marklogic-go/clients"
	"github.com/ryanjdew/go-marklogic-go/config"
	"github.com/ryanjdew/go-marklogic-go/documents"
	search "github.com/ryanjdew/go-marklogic-go/search"
	"github.com/ryanjdew/go-marklogic-go/semantics"
)

// Authentication options
const (
	BasicAuth  = clients.BasicAuth
	DigestAuth = clients.DigestAuth
	None       = clients.None
)

// Client is used for connecting to the MarkLogic REST API.
type Client clients.Client

type Connection clients.Connection

// NewClient creates the Client struct used for searching, etc.
func NewClient(conn *Connection) (*Client, error) {
	client, err := clients.NewClient(convertToClientConnection(conn))
	return convertToClient(client), err
}

// Config service
func (c *Client) Config() *config.Service {
	return config.NewService(convertToSubClient(c))
}

// Documents service
func (c *Client) Documents() *documents.Service {
	return documents.NewService(convertToSubClient(c))
}

// Search service
func (c *Client) Search() *search.Service {
	return search.NewService(convertToSubClient(c))
}

// Semantics service
func (c *Client) Semantics() *semantics.Service {
	return semantics.NewService(convertToSubClient(c))
}

func convertToSubClient(c *Client) *clients.Client {
	converted := clients.Client(*c)
	return &converted
}

func convertToClient(c *clients.Client) *Client {
	converted := Client(*c)
	return &converted
}

func convertToClientConnection(conn *Connection) *clients.Connection {
	converted := clients.Connection(*conn)
	return &converted
}
