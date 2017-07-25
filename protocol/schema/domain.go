// Code generated by cdpgen. DO NOT EDIT.

// Package schema implements the Schema domain. Provides information about the protocol schema.
package schema

import (
	"context"

	"github.com/mafredri/cdp/protocol/internal"
	"github.com/mafredri/cdp/rpcc"
)

// domainClient is a client for the Schema domain. Provides information about the protocol schema.
type domainClient struct{ conn *rpcc.Conn }

// NewClient returns a client for the Schema domain with the connection set to conn.
func NewClient(conn *rpcc.Conn) *domainClient {
	return &domainClient{conn: conn}
}

// GetDomains invokes the Schema method. Returns supported domains.
func (d *domainClient) GetDomains(ctx context.Context) (reply *GetDomainsReply, err error) {
	reply = new(GetDomainsReply)
	err = rpcc.Invoke(ctx, "Schema.getDomains", nil, reply, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "Schema", Op: "GetDomains", Err: err}
	}
	return
}