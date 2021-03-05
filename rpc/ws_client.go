package rpc

import "context"

type WsClient struct {
	idCounter     uint32
	reconnectFunc reconnectFunc
	writeConn     jsonWriter
}

func newWsClient(ctx context.Context, connect reconnectFunc) (*WsClient, error) {
	conn, err := connect(ctx)
	if err != nil {
		return nil, err
	}

	c := &WsClient{
		reconnectFunc: connect,
		writeConn:     conn,
	}
	return c, nil
}
