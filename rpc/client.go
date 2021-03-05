package rpc

import (
	"context"
	"fmt"
	"net/url"
)

type Client interface {
	Close()
}

func DialContext(ctx context.Context, rawrl string) (Client, error) {
	u, err := url.Parse(rawrl)
	if err != nil {
		return nil, err
	}
	switch u.Scheme {
	case "http", "https":
		return DialHTTP(rawrl)
	case "ws", "wss":
		return DialWebsocket(ctx, rawrl, "")
	default:
		return nil, fmt.Errorf("no known transport for URL scheme %q", u.Scheme)
	}
}
