package id

import (
	"errors"
	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
)

// Client is an ID client.
type Client struct {
	c *client.Client
}

// NewClient creates a new client.
func NewClient(c *client.Client) *Client {
	return &Client{c: c}
}

// SupportID checks if the server supports the ID extension.
func (c *Client) SupportID() (bool, error) {
	return c.c.Support(Capability)
}

// ID sends an ID command to the server and returns the server's ID.
func (c *Client) ID(clientID ID) (serverID ID, err error) {
	if state := c.c.State(); imap.ConnectedState&state != state {
		return nil, errors.New("Not connected")
	}

	var cmd imap.Commander = &Command{ID: clientID}

	res := &Response{}
	status, err := c.c.Execute(cmd, res)
	if err != nil {
		return
	}
	if err = status.Err(); err != nil {
		return
	}

	serverID = res.ID

	return
}
