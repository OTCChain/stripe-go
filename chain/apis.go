package chain

import (
	"encoding/json"
	"github.com/otcChain/chord-go/common"
	"github.com/otcChain/chord-go/rpc"
	"github.com/otcChain/chord-go/utils"
	"math/big"
)

func (cn *ChordNodeV1) initRpcApis() {
	rpc.HttpRpcApis["/tx/count"] = cn.AccountNonce
	rpc.HttpRpcApis["/chain/ID"] = cn.ChainID
	rpc.HttpRpcApis["/chain/Height"] = cn.ChainHeight
}

func (cn *ChordNodeV1) AccountNonce(request *rpc.JsonRpcMessageItem) (json.RawMessage, *rpc.JsonError) {
	var args []interface{}
	if err := json.Unmarshal(request.Params, &args); err != nil {
		return nil, &rpc.JsonError{Code: -1, Message: "invalid request parameters"}
	}
	if len(args) != 2 {
		return nil, &rpc.JsonError{
			Code:    -1,
			Message: "account nonce query need 2 parameters",
		}
	}
	addr, ok1 := args[0].(common.Address)
	status, ok2 := args[1].(string)
	if !(ok1 && ok2) {
		return nil, &rpc.JsonError{Code: -1, Message: "cast parameters failed"}
	}
	utils.LogInst().Debug().Msgf("account nonce query[%s-%s]", addr, status)
	return big.NewInt(1).Bytes(), nil
}

func (cn *ChordNodeV1) ChainHeight(*rpc.JsonRpcMessageItem) (json.RawMessage, *rpc.JsonError) {
	return nil, nil
}

func (cn *ChordNodeV1) ChainID(*rpc.JsonRpcMessageItem) (json.RawMessage, *rpc.JsonError) {
	return _nodeConfig.ChainID.Bytes(), nil
}
