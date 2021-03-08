package types

import (
	"github.com/herumi/bls-eth-go-binary/bls"
	"github.com/otcChain/chord-go/common"
	"math/big"
)

type ValTxData struct {
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
