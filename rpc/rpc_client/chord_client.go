package chordclient

import (
	"context"
	"github.com/otcChain/chord-go/rpc"
)

type Client struct {
	c rpc.Client
}

func Dial(rawUrl string) (*Client, error) {
	return DialContext(context.Background(), rawUrl)
}

func DialContext(ctx context.Context, rawUrl string) (*Client, error) {
	c, err := rpc.DialContext(ctx, rawUrl)
	if err != nil {
		return nil, err
	}
	return NewClient(c), nil
}

// NewClient creates a client that uses the given RPC client.
func NewClient(c rpc.Client) *Client {
	return &Client{c}
}

func (ec *Client) Close() {
	ec.c.Close()
}
