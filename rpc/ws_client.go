package rpc

import (
	"context"
	"encoding/base64"
	"github.com/gorilla/websocket"
	"net/http"
	"net/url"
	"sync"
)

const (
	wsReadBuffer  = 1024
	wsWriteBuffer = 1024
)

type WsClient struct {
	ctx      context.Context
	close    chan struct{}
	closing  chan struct{} // closed when client is quitting
	didClose chan struct{} // closed when client quits
}

func (c *WsClient) Call(result interface{}, url string, args ...interface{}) error {
	panic("implement me")
}

func (c *WsClient) CallContext(ctx context.Context, result interface{}, url string, args ...interface{}) error {
	panic("implement me")
}

var wsBufferPool = new(sync.Pool)

func DialWebsocket(ctx context.Context, endpoint, origin string) (Client, error) {
	dialer := websocket.Dialer{
		ReadBufferSize:  wsReadBuffer,
		WriteBufferSize: wsWriteBuffer,
		WriteBufferPool: wsBufferPool,
	}
	return DialWebsocketWithDialer(ctx, endpoint, origin, dialer)
}

func wsClientHeaders(endpoint, origin string) (string, http.Header, error) {
	endpointURL, err := url.Parse(endpoint)
	if err != nil {
		return endpoint, nil, err
	}
	header := make(http.Header)
	if origin != "" {
		header.Add("origin", origin)
	}
	if endpointURL.User != nil {
		b64auth := base64.StdEncoding.EncodeToString([]byte(endpointURL.User.String()))
		header.Add("authorization", "Basic "+b64auth)
		endpointURL.User = nil
	}
	return endpointURL.String(), header, nil
}

type wsHandshakeError struct {
	err    error
	status string
}

func (e wsHandshakeError) Error() string {
	s := e.err.Error()
	if e.status != "" {
		s += " (HTTP status " + e.status + ")"
	}
	return s
}

func DialWebsocketWithDialer(ctx context.Context,
	endpoint, origin string,
	dialer websocket.Dialer) (Client, error) {
	//endpoint, header, err := wsClientHeaders(endpoint, origin)
	//if err != nil {
	//	return nil, err
	//}
	//
	//conn, resp, err := dialer.DialContext(ctx, endpoint, header)

	return &WsClient{
		ctx:      ctx,
		close:    make(chan struct{}),
		closing:  make(chan struct{}),
		didClose: make(chan struct{}),
	}, nil
}

func (c *WsClient) Close() {
	select {
	case c.close <- struct{}{}:
		<-c.didClose
	case <-c.didClose:
	}
}
