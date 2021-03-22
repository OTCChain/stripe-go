package chordclient

import (
	"context"
	"encoding/json"
	"github.com/otcChain/chord-go/chord/transaction"
	"github.com/otcChain/chord-go/common"
	"github.com/otcChain/chord-go/utils/rlp"
)

type rpcTransaction struct {
	tx *transaction.Transaction
	txExtraInfo
}

type txExtraInfo struct {
	BlockNumber *string         `json:"blockNumber,omitempty"`
	BlockHash   *common.Hash    `json:"blockHash,omitempty"`
	From        *common.Address `json:"from,omitempty"`
}

func (tx *rpcTransaction) UnmarshalJSON(msg []byte) error {
	if err := json.Unmarshal(msg, &tx.tx); err != nil {
		return err
	}
	return json.Unmarshal(msg, &tx.txExtraInfo)
}

// SendTransaction injects a signed transaction into the pending pool for execution.
//
// If the transaction was a contract creation use the TransactionReceipt method to get the
// contract address after the transaction has been mined.
func (ec *Client) SendTransaction(ctx context.Context, tx transaction.Transaction) error {
	data, err := rlp.EncodeToBytes(tx)
	if err != nil {
		return err
	}
	_, err = ec.c.CallContext(ctx, "/tx/valTx", data)
	if err != nil {
		return err
	}
	return nil
}

// TransactionByHash returns the transaction with the given hash.
//func (ec *Client) TransactionByHash(ctx context.Context, hash common.Hash) (transaction *types.Transaction, isPending bool, err error) {
//	var json *rpcTransaction
//	err = ec.c.CallContext(ctx, &json, "eth_getTransactionByHash", hash)
//	if err != nil {
//		return nil, false, err
//	} else if json == nil {
//		return nil, false, ethereum.NotFound
//	} else if _, r, _ := json.transaction.RawSignatureValues(); r == nil {
//		return nil, false, fmt.Errorf("server returned transaction without signature")
//	}
//	if json.From != nil && json.BlockHash != nil {
//		setSenderFromServer(json.transaction, *json.From, *json.BlockHash)
//	}
//	return json.transaction, json.BlockNumber == nil, nil
//}
//
// TransactionSender returns the sender address of the given transaction. The transaction
// must be known to the remote node and included in the blockchain at the given block and
// index. The sender is the one derived by the protocol at the time of inclusion.
//
// There is a fast-path for transactions retrieved by TransactionByHash and
// TransactionInBlock. Getting their sender address can be done without an RPC interaction.
/*
func (ec *Client) TransactionSender(ctx context.Context,
				transaction *types.Transaction,
				block common.Hash,
				index uint) (common.Address, error) {
	// Try to load the address from the cache.
	sender, err := types.Sender(&senderFromServer{blockhash: block}, transaction)
	if err == nil {
		return sender, nil
	}
	var meta struct {
		Hash common.Hash
		From common.Address
	}
	if err = ec.c.CallContext(ctx, &meta, "eth_getTransactionByBlockHashAndIndex", block, hexutil.Uint64(index)); err != nil {
		return common.Address{}, err
	}
	if meta.Hash == (common.Hash{}) || meta.Hash != transaction.Hash() {
		return common.Address{}, fmt.Errorf("wrong inclusion block/index")
	}
	return meta.From, nil
}
*/

//// TransactionCount returns the total number of transactions in the given block.
//func (ec *Client) TransactionCount(ctx context.Context, blockHash common.Hash) (uint, error) {
//	var num hexutil.Uint
//	err := ec.c.CallContext(ctx, &num, "eth_getBlockTransactionCountByHash", blockHash)
//	return uint(num), err
//}
//
//// TransactionInBlock returns a single transaction at index in the given block.
//func (ec *Client) TransactionInBlock(ctx context.Context, blockHash common.Hash, index uint) (*types.Transaction, error) {
//	var json *rpcTransaction
//	err := ec.c.CallContext(ctx, &json, "eth_getTransactionByBlockHashAndIndex", blockHash, hexutil.Uint64(index))
//	if err != nil {
//		return nil, err
//	}
//	if json == nil {
//		return nil, ethereum.NotFound
//	} else if _, r, _ := json.transaction.RawSignatureValues(); r == nil {
//		return nil, fmt.Errorf("server returned transaction without signature")
//	}
//	if json.From != nil && json.BlockHash != nil {
//		setSenderFromServer(json.transaction, *json.From, *json.BlockHash)
//	}
//	return json.transaction, err
//}
//
//// TransactionReceipt returns the receipt of a transaction by transaction hash.
//// Note that the receipt is not available for pending transactions.
//func (ec *Client) TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
//	var r *types.Receipt
//	err := ec.c.CallContext(ctx, &r, "eth_getTransactionReceipt", txHash)
//	if err == nil {
//		if r == nil {
//			return nil, ethereum.NotFound
//		}
//	}
//	return r, err
//}
//
//func toBlockNumArg(number *big.Int) string {
//	if number == nil {
//		return "latest"
//	}
//	pending := big.NewInt(-1)
//	if number.Cmp(pending) == 0 {
//		return "pending"
//	}
//	return hexutil.EncodeBig(number)
//}
