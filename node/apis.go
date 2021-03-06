package node

import (
	"encoding/json"
	"github.com/otcChain/chord-go/rpc"
)

func (cn *ChordNodeV1) initRpcApis() {
	rpc.HttpRpcApis["/chain/ID"] = cn.ChainID
	rpc.HttpRpcApis["/chain/Height"] = cn.ChainHeight
}

func (cn *ChordNodeV1) ChainID(*rpc.JsonRpcMessageItem) (json.RawMessage, *rpc.JsonError) {
	return nil, nil
}
func (cn *ChordNodeV1) ChainHeight(*rpc.JsonRpcMessageItem) (json.RawMessage, *rpc.JsonError) {
	return nil, nil
}
