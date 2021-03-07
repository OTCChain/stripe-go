package rpc

type WsRpc struct {
}

func (wr *WsRpc) StartRpc() chan error {
	return nil
}

func newWsRpc() *WsRpc {
	wr := &WsRpc{}
	return wr
}
