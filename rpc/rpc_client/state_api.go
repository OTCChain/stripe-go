package chordclient

import (
	"context"
	"fmt"
	"math/big"
)

// State Access
// NetworkID returns the network ID (also known as the chain ID) for this chain.
func (ec *Client) NetworkID(ctx context.Context) (*big.Int, error) {
	version := new(big.Int)
	var ver string
	if err := ec.c.CallContext(ctx, &ver, "/p2p/nid"); err != nil {
		return nil, err
	}
	if _, ok := version.SetString(ver, 10); !ok {
		return nil, fmt.Errorf("invalid net_version result %q", ver)
	}
	return version, nil
}

//
//// BalanceAt returns the wei balance of the given account.
//// The block number can be nil, in which case the balance is taken from the latest known block.
//func (ec *Client) BalanceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (*big.Int, error) {
//	var result hexutil.Big
//	err := ec.c.CallContext(ctx, &result, "eth_getBalance", account, toBlockNumArg(blockNumber))
//	return (*big.Int)(&result), err
//}
//
//// StorageAt returns the value of key in the contract storage of the given account.
//// The block number can be nil, in which case the value is taken from the latest known block.
//func (ec *Client) StorageAt(ctx context.Context, account common.Address, key common.Hash, blockNumber *big.Int) ([]byte, error) {
//	var result hexutil.Bytes
//	err := ec.c.CallContext(ctx, &result, "eth_getStorageAt", account, key, toBlockNumArg(blockNumber))
//	return result, err
//}
//
//// CodeAt returns the contract code of the given account.
//// The block number can be nil, in which case the code is taken from the latest known block.
//func (ec *Client) CodeAt(ctx context.Context, account common.Address, blockNumber *big.Int) ([]byte, error) {
//	var result hexutil.Bytes
//	err := ec.c.CallContext(ctx, &result, "eth_getCode", account, toBlockNumArg(blockNumber))
//	return result, err
//}
//
//// NonceAt returns the account nonce of the given account.
//// The block number can be nil, in which case the nonce is taken from the latest known block.
//func (ec *Client) NonceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (uint64, error) {
//	var result hexutil.Uint64
//	err := ec.c.CallContext(ctx, &result, "eth_getTransactionCount", account, toBlockNumArg(blockNumber))
//	return uint64(result), err
//}
