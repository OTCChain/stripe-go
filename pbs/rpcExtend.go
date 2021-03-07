package pbs

func (x *RpcResponse) RpcError(msg string) {
	x.Data = []byte(msg)
}

func RpcError(msg string) *RpcResponse {
	return &RpcResponse{
		Code: -1,
		Data: []byte(msg),
	}
}

func RpcOk(d []byte) (x *RpcResponse) {
	return &RpcResponse{
		Code: 0,
		Data: d,
	}
}
