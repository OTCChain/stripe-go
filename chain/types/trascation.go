package types

import (
	"github.com/herumi/bls-eth-go-binary/bls"
	"github.com/otcChain/chord-go/common"
	"github.com/otcChain/chord-go/utils/rlp"
	"math/big"
	"sync/atomic"
	"time"
)

type Transaction struct {
	inner TxData    // Consensus contents of a transaction
	time  time.Time // Time first seen locally (spam avoidance)

	// caches
	hash atomic.Value
	size atomic.Value
	from atomic.Value
}

// NewTx creates a new transaction.
func NewTx(inner TxData) *Transaction {
	tx := new(Transaction)
	tx.inner = inner
	tx.time = time.Now()
	tx.size.Store(0)
	return tx
}

type TxData struct {
	Nonce   uint64   `json:"nonce"      gencodec:"required"`
	Price   *big.Int `json:"gasPrice"   gencodec:"required"`
	Gas     uint64   `json:"gas"        gencodec:"required"`
	ShardID uint32   `json:"shardID"    gencodec:"required"`
	//ToShardID    uint32          `json:"toShardID"  gencodec:"required"`
	To    *common.Address `json:"to"         rlp:"nil"` // nil means contract creation
	Value *big.Int        `json:"value"      gencodec:"required"`
	Data  []byte          `json:"input"      gencodec:"required"`

	Hash common.Hash `json:"hash" rlp:"-"`
	Sig  *bls.Sign   `json:"signature"      gencodec:"required"`
}

func (tx *Transaction) MarshalBinary() ([]byte, error) {
	return rlp.EncodeToBytes(tx.inner)
}

func (tx *Transaction) SignTx(prv *bls.SecretKey) error {
	txBytes, err := tx.MarshalBinary()
	if err != nil {
		return err
	}
	tx.inner.Hash = common.BytesToHash(txBytes)
	tx.inner.Sig = prv.SignByte(txBytes)
	return nil
}

func (tx *Transaction) Hash() common.Hash {
	return tx.inner.Hash
}
