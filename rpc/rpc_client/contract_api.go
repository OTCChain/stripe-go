package chordclient

// Contract Calling

// CallContract executes a message call transaction, which is directly executed in the VM
// of the node, but never mined into the blockchain.
//
// blockNumber selects the block height at which the call runs. It can be nil, in which
// case the code is taken from the latest known block. Note that state from very old
// blocks might not be available.
//func (ec *Client) CallContract(ctx context.Context, msg ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
//	var hex hexutil.Bytes
//	err := ec.c.CallContext(ctx, &hex, "eth_call", toCallArg(msg), toBlockNumArg(blockNumber))
//	if err != nil {
//		return nil, err
//	}
//	return hex, nil
//}
//
//// PendingCallContract executes a message call transaction using the EVM.
//// The state seen by the contract call is the pending state.
//func (ec *Client) PendingCallContract(ctx context.Context, msg ethereum.CallMsg) ([]byte, error) {
//	var hex hexutil.Bytes
//	err := ec.c.CallContext(ctx, &hex, "eth_call", toCallArg(msg), "pending")
//	if err != nil {
//		return nil, err
//	}
//	return hex, nil
//}
//
//// SuggestGasPrice retrieves the currently suggested gas price to allow a timely
//// execution of a transaction.
//func (ec *Client) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
//	var hex hexutil.Big
//	if err := ec.c.CallContext(ctx, &hex, "eth_gasPrice"); err != nil {
//		return nil, err
//	}
//	return (*big.Int)(&hex), nil
//}
//
//// EstimateGas tries to estimate the gas needed to execute a specific transaction based on
//// the current pending state of the backend blockchain. There is no guarantee that this is
//// the true gas limit requirement as other transactions may be added or removed by miners,
//// but it should provide a basis for setting a reasonable default.
//func (ec *Client) EstimateGas(ctx context.Context, msg ethereum.CallMsg) (uint64, error) {
//	var hex hexutil.Uint64
//	err := ec.c.CallContext(ctx, &hex, "eth_estimateGas", toCallArg(msg))
//	if err != nil {
//		return 0, err
//	}
//	return uint64(hex), nil
//}
//

//func toCallArg(msg ethereum.CallMsg) interface{} {
//	arg := map[string]interface{}{
//		"from": msg.From,
//		"to":   msg.To,
//	}
//	if len(msg.Data) > 0 {
//		arg["data"] = hexutil.Bytes(msg.Data)
//	}
//	if msg.Value != nil {
//		arg["value"] = (*hexutil.Big)(msg.Value)
//	}
//	if msg.Gas != 0 {
//		arg["gas"] = hexutil.Uint64(msg.Gas)
//	}
//	if msg.GasPrice != nil {
//		arg["gasPrice"] = (*hexutil.Big)(msg.GasPrice)
//	}
//	return arg
//}
