package chordclient

// Blockchain Access

// ChainId retrieves the current chain ID for transaction replay protection.
//func (ec *Client) ChainID(ctx context.Context) (*big.Int, error) {
//	var result hexutil.Big
//	err := ec.c.CallContext(ctx, &result, "eth_chainId")
//	if err != nil {
//		return nil, err
//	}
//	return (*big.Int)(&result), err
//}
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
