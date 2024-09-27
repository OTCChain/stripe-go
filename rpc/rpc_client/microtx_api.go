package chordclient

import (
	"context"
	"github.com/otcChain/chord-go/chord/transaction"
	"github.com/otcChain/chord-go/utils/rlp"
)

func (ec *Client) SendMicroTx(ctx context.Context, tx transaction.Transaction) error {
	data, err := rlp.EncodeToBytes(tx)
	if err != nil {
		return err
	}
	_, err = ec.c.CallContext(ctx, "/transaction/newMicroTx", data)
	if err != nil {
		return err
	}
	return nil
}
