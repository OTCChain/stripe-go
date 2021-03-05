package rpc

import (
	"context"
	//mapset "github.com/deckarep/golang-set"
	"net/http"
	"net/url"
)

type HttpClient struct {
	ctx context.Context
}

func DialHTTP(endpoint string) (Client, error) {
	return DialHTTPWithClient(endpoint, new(http.Client))
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
	return &HttpClient{
		ctx: ctx,
	}, nil
}

func (h HttpClient) Close() {
	return
}
