package rpc

type HttpRpc struct {
}

func (hr *HttpRpc) StartRpc() error {
	return nil
}

func newHttpRpc() Rpc {
	hr := &HttpRpc{}
	return hr
}
