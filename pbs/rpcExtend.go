package pbs

func (x *RpcResponse) RpcError(msg string) {
	x.Data = []byte(msg)
}

//TODO:: error code list
func RpcError(msg string) *RpcResponse {
	return &RpcResponse{
		Code: int32(ApiRet_Error),
		Data: []byte(msg),
	}
}

func RpcOk(d []byte) (x *RpcResponse) {
	return &RpcResponse{
		Code: int32(ApiRet_OK),
		Data: d,
	}
}
