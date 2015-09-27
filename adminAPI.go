package goMarklogicGo

import (
	admin "github.com/ryanjdew/go-marklogic-go/admin"
	clients "github.com/ryanjdew/go-marklogic-go/clients"
)

// AdminClient is used for connecting to the MarkLogic Management API.
type AdminClient clients.AdminClient

// NewAdminClient creates the Client struct used for managing databases, etc.
func NewAdminClient(conn *Connection) (*AdminClient, error) {
	client, err := clients.NewAdminClient(convertToClientConnection(conn))
	return convertToAdminClient(client), err
}

// Admin service
func (c *AdminClient) Service() *admin.Service {
	return admin.NewService(convertToSubAdminClient(c))
}

func convertToSubAdminClient(mc *AdminClient) *clients.AdminClient {
	converted := clients.AdminClient(*mc)
	return &converted
}

func convertToAdminClient(mc *clients.AdminClient) *AdminClient {
	converted := AdminClient(*mc)
	return &converted
}
