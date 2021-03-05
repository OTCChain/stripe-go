package rpc

import (
	"context"
	"net/http"
	"net/url"
)

type HttpClient struct {
	idCounter     uint32
	reconnectFunc reconnectFunc
	writeConn     jsonWriter
}
type reconnectFunc func(ctx context.Context) (ServerCodec, error)

func newHttpClient(ctx context.Context, connect reconnectFunc) (*HttpClient, error) {
	conn, err := connect(ctx)
	if err != nil {
		return nil, err
	}
	c := &HttpClient{
		reconnectFunc: connect,
		writeConn:     conn,
	}
	return c, nil
}

func DialHTTPWithClient(endpoint string, client *http.Client) (*HttpClient, error) {
	_, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	headers := make(http.Header, 2)
	headers.Set("accept", contentType)
	headers.Set("content-type", contentType)

	var reconnFunc = func(ctx context.Context) (ServerCodec, error) {
		hc := &httpConn{
			client:  client,
			headers: headers,
			url:     endpoint,
			closeCh: make(chan interface{}),
		}
		return hc, nil
	}

	return newHttpClient(ctx, reconnFunc)
}
