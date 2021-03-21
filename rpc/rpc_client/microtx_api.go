package chordclient

import (
	"context"
	"github.com/otcChain/chord-go/chord/types"
)

func (ec *Client) SendMicroTx(ctx context.Context, tx *types.Transaction) error {
	data, err := tx.MarshalBinary()
	if err != nil {
		return err
	}
	_, err = ec.c.CallContext(ctx, "/tx/newMicroTx", data)
	if err != nil {
		return err
	}
	return nil
}
