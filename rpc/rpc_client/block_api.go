package chordclient

import (
	"context"
	"math/big"
)

// Blockchain Access

//ChainId retrieves the current chain ID for transaction replay protection.
func (ec *Client) ChainID(ctx context.Context) (*big.Int, error) {
	var result big.Int
	data, err := ec.c.CallContext(ctx, "/chain/ID", nil)
	if err != nil {
		return nil, err
	}
	result.SetBytes(data)
	return &result, err
}

//
//// BlockByHash returns the given full block.
////
//// Note that loading full blocks requires two requests. Use HeaderByHash
//// if you don't need all transactions or uncle headers.
//func (ec *Client) BlockByHash(ctx context.Context, hash common.Hash) (*types.Block, error) {
//	return ec.getBlock(ctx, "eth_getBlockByHash", hash, true)
//}
//
//// BlockByNumber returns a block from the current canonical chain. If number is nil, the
//// latest known block is returned.
////
//// Note that loading full blocks requires two requests. Use HeaderByNumber
//// if you don't need all transactions or uncle headers.
//func (ec *Client) BlockByNumber(ctx context.Context, number *big.Int) (*types.Block, error) {
//	return ec.getBlock(ctx, "eth_getBlockByNumber", toBlockNumArg(number), true)
//}
//
//// BlockNumber returns the most recent block number
//func (ec *Client) BlockNumber(ctx context.Context) (uint64, error) {
//	var result hexutil.Uint64
//	err := ec.c.CallContext(ctx, &result, "eth_blockNumber")
//	return uint64(result), err
//}

//type rpcBlock struct {
//	Hash         common.Hash      `json:"hash"`
//	Transactions []rpcTransaction `json:"transactions"`
//	UncleHashes  []common.Hash    `json:"uncles"`
//}
//
//func (ec *Client) getBlock(ctx context.Context, method string, args ...interface{}) (*types.Block, error) {
//	var raw json.RawMessage
//	err := ec.c.CallContext(ctx, &raw, method, args...)
//	if err != nil {
//		return nil, err
//	} else if len(raw) == 0 {
//		return nil, ethereum.NotFound
//	}
//	// Decode header and transactions.
//	var head *types.Header
//	var body rpcBlock
//	if err := json.Unmarshal(raw, &head); err != nil {
//		return nil, err
//	}
//	if err := json.Unmarshal(raw, &body); err != nil {
//		return nil, err
//	}
//	// Quick-verify transaction and uncle lists. This mostly helps with debugging the server.
//	if head.UncleHash == types.EmptyUncleHash && len(body.UncleHashes) > 0 {
//		return nil, fmt.Errorf("server returned non-empty uncle list but block header indicates no uncles")
//	}
//	if head.UncleHash != types.EmptyUncleHash && len(body.UncleHashes) == 0 {
//		return nil, fmt.Errorf("server returned empty uncle list but block header indicates uncles")
//	}
//	if head.TxHash == types.EmptyRootHash && len(body.Transactions) > 0 {
//		return nil, fmt.Errorf("server returned non-empty transaction list but block header indicates no transactions")
//	}
//	if head.TxHash != types.EmptyRootHash && len(body.Transactions) == 0 {
//		return nil, fmt.Errorf("server returned empty transaction list but block header indicates transactions")
//	}
//	// Load uncles because they are not included in the block response.
//	var uncles []*types.Header
//	if len(body.UncleHashes) > 0 {
//		uncles = make([]*types.Header, len(body.UncleHashes))
//		reqs := make([]rpc.BatchElem, len(body.UncleHashes))
//		for i := range reqs {
//			reqs[i] = rpc.BatchElem{
//				Method: "eth_getUncleByBlockHashAndIndex",
//				Args:   []interface{}{body.Hash, hexutil.EncodeUint64(uint64(i))},
//				Result: &uncles[i],
//			}
//		}
//		if err := ec.c.BatchCallContext(ctx, reqs); err != nil {
//			return nil, err
//		}
//		for i := range reqs {
//			if reqs[i].Error != nil {
//				return nil, reqs[i].Error
//			}
//			if uncles[i] == nil {
//				return nil, fmt.Errorf("got null header for uncle %d of block %x", i, body.Hash[:])
//			}
//		}
//	}
//	// Fill the sender cache of transactions in the block.
//	txs := make([]*types.Transaction, len(body.Transactions))
//	for i, tx := range body.Transactions {
//		if tx.From != nil {
//			setSenderFromServer(tx.tx, *tx.From, body.Hash)
//		}
//		txs[i] = tx.tx
//	}
//	return types.NewBlockWithHeader(head).WithBody(txs, uncles), nil
//}
//
//// HeaderByHash returns the block header with the given hash.
//func (ec *Client) HeaderByHash(ctx context.Context, hash common.Hash) (*types.Header, error) {
//	var head *types.Header
//	err := ec.c.CallContext(ctx, &head, "eth_getBlockByHash", hash, false)
//	if err == nil && head == nil {
//		err = ethereum.NotFound
//	}
//	return head, err
//}
//
//// HeaderByNumber returns a block header from the current canonical chain. If number is
//// nil, the latest known header is returned.
//func (ec *Client) HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
//	var head *types.Header
//	err := ec.c.CallContext(ctx, &head, "eth_getBlockByNumber", toBlockNumArg(number), false)
//	if err == nil && head == nil {
//		err = ethereum.NotFound
//	}
//	return head, err
//}
