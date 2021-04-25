package chordclient

import (
	"context"
	"github.com/golang/protobuf/proto"
	"github.com/otcChain/chord-go/common"
	pbs "github.com/otcChain/chord-go/pbs/rpc"
	"math/big"
)

// Pending State

// PendingBalanceAt returns the wei balance of the given account in the pending state.
//func (ec *Client) PendingBalanceAt(ctx context.Context, account common.Address) (*big.Int, error) {
//	var result hexutil.Big
//	err := ec.c.CallContext(ctx, &result, "eth_getBalance", account, "pending")
//	return (*big.Int)(&result), err
//}
//
//// PendingStorageAt returns the value of key in the contract storage of the given account in the pending state.
//func (ec *Client) PendingStorageAt(ctx context.Context, account common.Address, key common.Hash) ([]byte, error) {
//	var result hexutil.Bytes
//	err := ec.c.CallContext(ctx, &result, "eth_getStorageAt", account, key, "pending")
//	return result, err
//}
//
//// PendingCodeAt returns the contract code of the given account in the pending state.
//func (ec *Client) PendingCodeAt(ctx context.Context, account common.Address) ([]byte, error) {
//	var result hexutil.Bytes
//	err := ec.c.CallContext(ctx, &result, "eth_getCode", account, "pending")
//	return result, err
//}
//
// PendingNonceAt returns the account nonce of the given account in the pending state.
// This is the nonce that should be used for the next transaction.

func (ec *Client) PendingNonceAt(ctx context.Context, account common.Address) (uint64, error) {
	var result big.Int
	param := &pbs.AccountNonce{
		Account: account.String(),
		Status:  pbs.TxType_Pending,
	}
	protoData, err := proto.Marshal(param)
	if err != nil {
		return 0, err
	}
	byts, err := ec.c.CallContext(ctx, "/account/nonce", protoData)
	if err != nil {
		return 0, err
	}
	result.SetBytes(byts)
	return result.Uint64(), nil
}

//// PendingTransactionCount returns the total number of transactions in the pending state.
//func (ec *Client) PendingTransactionCount(ctx context.Context) (uint, error) {
//	var num hexutil.Uint
//	err := ec.c.CallContext(ctx, &num, "eth_getBlockTransactionCountByNumber", "pending")
//	return uint(num), err
//}
//
//// TODO: SubscribePendingTransactions (needs server side)
