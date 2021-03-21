package types

import (
	"crypto"
	"github.com/otcChain/chord-go/common"
	"github.com/otcChain/chord-go/utils/rlp"
	"time"
)

type TxData interface {
	Hash() common.Hash
	Sig() crypto.PrivateKey
	SignTx(crypto.PrivateKey) error
}

type Transaction struct {
	inner TxData    // Consensus contents of a transaction
	time  time.Time // Time first seen locally (spam avoidance)
}

// NewTx creates a new transaction.
func NewTx(inner TxData) *Transaction {
	tx := new(Transaction)
	tx.inner = inner
	tx.time = time.Now()
	return tx
}

func (tx *Transaction) MarshalBinary() ([]byte, error) {
	return rlp.EncodeToBytes(tx.inner)
}

func (tx *Transaction) SignTx(pri crypto.PrivateKey) error {
	return tx.inner.SignTx(pri)
}

func (tx *Transaction) Hash() common.Hash {
	return tx.inner.Hash()
}
