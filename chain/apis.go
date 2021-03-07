package chain

import (
	"encoding/json"
	"github.com/otcChain/chord-go/common"
	"github.com/otcChain/chord-go/rpc"
	"github.com/otcChain/chord-go/utils"
	"math/big"
	"reflect"
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
		utils.LogInst().Warn().Msgf("account nonce query args [%d](%s)", len(args), string(request.Params))
		return nil, rpc.JSError("account nonce query need 2 parameters")
	}

	addr, err := common.InterfaceToAddress(args[0].(string))
	if err != nil {
		return nil, rpc.JError(err)
	}
	status, ok2 := args[1].(string)
	if !ok2 {
		utils.LogInst().Debug().
			Bool("cast arg1 ret", ok2).
			Str("addr type:", reflect.TypeOf(args[0]).String()).Send()
		return nil, &rpc.JsonError{Code: -1, Message: "cast parameters failed"}
	}
	utils.LogInst().Debug().Str("account", addr.String()).
		Str("tx status:", status).
		Msg("account nonce query")
	return big.NewInt(1).Bytes(), nil
}

func (cn *ChordNodeV1) ChainHeight(*rpc.JsonRpcMessageItem) (json.RawMessage, *rpc.JsonError) {
	return nil, nil
}

func (cn *ChordNodeV1) ChainID(*rpc.JsonRpcMessageItem) (json.RawMessage, *rpc.JsonError) {
	return _nodeConfig.ChainID.Bytes(), nil
}
