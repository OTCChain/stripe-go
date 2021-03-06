package chordclient

// Filters

// FilterLogs executes a filter query.
//func (ec *Client) FilterLogs(ctx context.Context, q ethereum.FilterQuery) ([]types.Log, error) {
//	var result []types.Log
//	arg, err := toFilterArg(q)
//	if err != nil {
//		return nil, err
//	}
//	err = ec.c.CallContext(ctx, &result, "eth_getLogs", arg)
//	return result, err
//}
//
//// SubscribeFilterLogs subscribes to the results of a streaming filter query.
//func (ec *Client) SubscribeFilterLogs(ctx context.Context, q ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error) {
//	arg, err := toFilterArg(q)
//	if err != nil {
//		return nil, err
//	}
//	return ec.c.EthSubscribe(ctx, ch, "logs", arg)
//}
//
//func toFilterArg(q ethereum.FilterQuery) (interface{}, error) {
//	arg := map[string]interface{}{
//		"address": q.Addresses,
//		"topics":  q.Topics,
//	}
//	if q.BlockHash != nil {
//		arg["blockHash"] = *q.BlockHash
//		if q.FromBlock != nil || q.ToBlock != nil {
//			return nil, fmt.Errorf("cannot specify both BlockHash and FromBlock/ToBlock")
//		}
//	} else {
//		if q.FromBlock == nil {
//			arg["fromBlock"] = "0x0"
//		} else {
//			arg["fromBlock"] = toBlockNumArg(q.FromBlock)
//		}
//		arg["toBlock"] = toBlockNumArg(q.ToBlock)
//	}
//	return arg, nil
//}
