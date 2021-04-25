package chord

import (
	pbs "github.com/otcChain/chord-go/pbs/rpc"
	"github.com/otcChain/chord-go/rpc"
	"github.com/otcChain/chord-go/utils"
	"google.golang.org/protobuf/proto"
	"math/big"
)

func (cn *NodeV1) initRpcApis() {
	rpc.HttpRpcApis["/account/nonce"] = cn.AccountNonce
	rpc.HttpRpcApis["/tx/microTx"] = cn.NewMicroTx
	rpc.HttpRpcApis["/tx/valTx"] = cn.NewValTx
	rpc.HttpRpcApis["/chord/ID"] = cn.ChainID
	rpc.HttpRpcApis["/chord/Height"] = cn.ChainHeight
}

// ChainHeight ---------------------------Chord API-------------------------
func (cn *NodeV1) ChainHeight(_ *pbs.RpcMsgItem) *pbs.RpcResponse {
	return pbs.RpcOk(nil)
}

func (cn *NodeV1) ChainID(_ *pbs.RpcMsgItem) *pbs.RpcResponse {
	return pbs.RpcOk(_nodeConfig.ChainID.Bytes())
}

// AccountNonce ---------------------------Account API-------------------------
func (cn *NodeV1) AccountNonce(request *pbs.RpcMsgItem) *pbs.RpcResponse {
	var args = &pbs.AccountNonce{}
	if err := proto.Unmarshal(request.Parameter, args); err != nil {
		return pbs.RpcError("invalid request parameters")
	}
	utils.LogInst().Debug().Str("account", args.Account).
		Str("transaction status:", args.Status.String()).
		Msg("account nonce query")
	return pbs.RpcOk(big.NewInt(12345679890).Bytes())
}

// NewMicroTx ---------------------------TX API-------------------------
func (cn *NodeV1) NewMicroTx(_ *pbs.RpcMsgItem) *pbs.RpcResponse {
	return pbs.RpcOk(nil)
}

func (cn *NodeV1) NewValTx(_ *pbs.RpcMsgItem) *pbs.RpcResponse {
	return pbs.RpcOk(nil)
}
