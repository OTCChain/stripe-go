package transaction

import (
	"crypto"
	"github.com/otcChain/chord-go/common"
	"math/big"
	"time"
)

type Transaction interface {
	Hash() common.Hash
	SignTx(crypto.PrivateKey) error
}

func NewValTx(nonce uint64,
	receipt *common.Address,
	val, price *big.Int,
	chID uint32,
	gas uint64,
	data []byte) Transaction {

	vt := ValTx{
		valTxData: &valTxData{
			Nonce:   nonce,
			Price:   price,
			Gas:     gas,
			ChainID: chID,
			To:      receipt,
			Value:   val,
			Data:    data,
		},
		time: time.Now(),
	}

	return vt
}

func NewMicTx(nonce uint64,
	receipt *common.Address,
	val *big.Int,
	chID uint32) Transaction {

	mt := &MicroTx{
		microTxData: &microTxData{
			Nonce:   nonce,
			To:      receipt,
			Value:   val,
			ChainID: chID,
		},
		time: time.Now(),
	}
	return mt
}
