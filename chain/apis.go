package chain

import (
	"github.com/golang/protobuf/proto"
	"github.com/otcChain/chord-go/common"
	"github.com/otcChain/chord-go/pbs"
	"github.com/otcChain/chord-go/rpc"
	"github.com/otcChain/chord-go/utils"
	"math/big"
)

func (cn *ChordNodeV1) initRpcApis() {
	rpc.HttpRpcApis["/account/nonce"] = cn.AccountNonce
	rpc.HttpRpcApis["/chain/ID"] = cn.ChainID
	rpc.HttpRpcApis["/chain/Height"] = cn.ChainHeight
}

type APIAccountNonce struct {
	Account common.Address
	Status  string
}

func (cn *ChordNodeV1) AccountNonce(request *pbs.RpcMsgItem) *pbs.RpcResponse {
	var args = &pbs.AccountNonce{}
	if err := proto.Unmarshal(request.Parameter, args); err != nil {
		return pbs.RpcError("invalid request parameters")
	}
	utils.LogInst().Debug().Str("account", args.Account).
		Str("tx status:", args.Status.String()).
		Msg("account nonce query")
	return pbs.RpcOk(big.NewInt(12345679890).Bytes())
}

func (cn *ChordNodeV1) ChainHeight(_ *pbs.RpcMsgItem) *pbs.RpcResponse {
	return pbs.RpcOk(nil)
}

func (cn *ChordNodeV1) ChainID(_ *pbs.RpcMsgItem) *pbs.RpcResponse {
	return pbs.RpcOk(_nodeConfig.ChainID.Bytes())
}
