package rpc

type WsRpc struct {
}

func (wr *WsRpc) StartRpc() error {
	return nil
}

func newWsRpc() Rpc {
	wr := &WsRpc{}
	return wr
}
