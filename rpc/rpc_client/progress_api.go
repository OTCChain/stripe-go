package chordclient

//type rpcProgress struct {
//	StartingBlock hexutil.Uint64
//	CurrentBlock  hexutil.Uint64
//	HighestBlock  hexutil.Uint64
//	PulledStates  hexutil.Uint64
//	KnownStates   hexutil.Uint64
//}
//
//// SyncProgress retrieves the current progress of the sync algorithm. If there's
//// no sync currently running, it returns nil.
//func (ec *Client) SyncProgress(ctx context.Context) (*ethereum.SyncProgress, error) {
//	var raw json.RawMessage
//	if err := ec.c.CallContext(ctx, &raw, "eth_syncing"); err != nil {
//		return nil, err
//	}
//	// Handle the possible response types
//	var syncing bool
//	if err := json.Unmarshal(raw, &syncing); err == nil {
//		return nil, nil // Not syncing (always false)
//	}
//	var progress *rpcProgress
//	if err := json.Unmarshal(raw, &progress); err != nil {
//		return nil, err
//	}
//	return &ethereum.SyncProgress{
//		StartingBlock: uint64(progress.StartingBlock),
//		CurrentBlock:  uint64(progress.CurrentBlock),
//		HighestBlock:  uint64(progress.HighestBlock),
//		PulledStates:  uint64(progress.PulledStates),
//		KnownStates:   uint64(progress.KnownStates),
//	}, nil
//}
//
//// SubscribeNewHead subscribes to notifications about the current blockchain head
//// on the given channel.
//func (ec *Client) SubscribeNewHead(ctx context.Context, ch chan<- *types.Header) (ethereum.Subscription, error) {
//	return ec.c.EthSubscribe(ctx, ch, "newHeads")
//}
