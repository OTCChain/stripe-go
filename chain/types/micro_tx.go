package types

import (
	"crypto"
	"fmt"
	"github.com/herumi/bls-eth-go-binary/bls"
	"github.com/otcChain/chord-go/common"
	"github.com/otcChain/chord-go/utils/rlp"
	"math/big"
)

type MicroTxData struct {
	Nonce uint64         `json:"nonce"      gencodec:"required"`
	From  common.Address `json:"from"      gencodec:"required"`
	To    common.Address `json:"to"      gencodec:"required"`
	Value *big.Int       `json:"value"      gencodec:"required"`
	Gas   uint64         `json:"gas"        gencodec:"required"`
	Sig   *bls.Sign      `json:"sig" rlp:"-"`
}

func (m MicroTxData) Sign(pri crypto.PrivateKey) error {
	prv, ok := pri.(*bls.SecretKey)
	if !ok {
		return fmt.Errorf("invalid micro transaction private key for singer")
	}
	bts, err := rlp.EncodeToBytes(m)
	if err != nil {
		return err
	}
	m.Sig = prv.SignByte(bts)
	return nil
}
