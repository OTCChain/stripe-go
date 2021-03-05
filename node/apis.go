package node

import (
	"encoding/json"
	"github.com/otcChain/chord-go/rpc"
)

var (
	publicHttpRpcApi = make(rpc.HttpApiRouter)
)

func (cn *ChordNodeV1) initRouter() {
	publicHttpRpcApi["/chain/ID"] = cn.ChainID
	publicHttpRpcApi["/chain/Height"] = cn.ChainHeight
}

func (cn *ChordNodeV1) ChainID(*rpc.JsonRpcMessageItem) (json.RawMessage, *rpc.JsonError) {
	return nil, nil
}
func (cn *ChordNodeV1) ChainHeight(*rpc.JsonRpcMessageItem) (json.RawMessage, *rpc.JsonError) {
	return nil, nil
}
