package transaction

import (
	"crypto"
	"fmt"
	"github.com/herumi/bls-eth-go-binary/bls"
	"github.com/otcChain/chord-go/common"
	"github.com/otcChain/chord-go/utils/rlp"
	"math/big"
	"sync/atomic"
	"time"
)

type microTxData struct {
	Nonce     uint64          `json:"nonce"      gencodec:"required"`
	To        *common.Address `json:"to"      gencodec:"required"`
	Value     *big.Int        `json:"value"      gencodec:"required"`
	ChainID   uint32          `json:"chainIDs"      gencodec:"required"`
	ShardID   uint32          `json:"shardID"    gencodec:"required"`
	ToShardID uint32          `json:"toShardID"  gencodec:"required"`
}

type MicroTx struct {
	*microTxData
	from      atomic.Value
	time      time.Time
	hash      common.Hash
	Signature *bls.Sign `json:"sig" gencodec:"required"`
}

func (mt *MicroTx) Hash() common.Hash {
	return mt.hash
}

func (mt *MicroTx) SignTx(pri crypto.PrivateKey) error {
	prv, ok := pri.(*bls.SecretKey)
	if !ok {
		return fmt.Errorf("invalid micro transaction private key for singer")
	}
	bts, err := rlp.EncodeToBytes(mt.microTxData)
	if err != nil {
		return err
	}
	mt.Signature = prv.SignByte(bts)
	return nil
}
