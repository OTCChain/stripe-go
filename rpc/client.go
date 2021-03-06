package rpc

import (
	"context"
	"fmt"
	"net/url"
)

type Client interface {
	Close()
	Call(result interface{}, url string, args ...interface{}) error
	CallContext(ctx context.Context, result interface{}, path string, args ...interface{}) error
}

func DialContext(ctx context.Context, rawUrl string) (Client, error) {
	u, err := url.Parse(rawUrl)
	if err != nil {
		return nil, err
	}
	switch u.Scheme {
	case "http", "https":
		return DialHTTP(rawUrl)
	case "ws", "wss":
		return DialWebsocket(ctx, rawUrl, "")
	default:
		return nil, fmt.Errorf("no known transport for URL scheme %q", u.Scheme)
	}
}
