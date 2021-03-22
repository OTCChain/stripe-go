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

type valTxData struct {
	Nonce   uint64          `json:"nonce"      gencodec:"required"`
	Price   *big.Int        `json:"gasPrice"   gencodec:"required"`
	Gas     uint64          `json:"gas"        gencodec:"required"`
	ChainID uint32          `json:"chainIDs"      gencodec:"required"`
	To      *common.Address `json:"to"         rlp:"nil"` // nil means contract creation
	Value   *big.Int        `json:"value"      gencodec:"required"`
	Data    []byte          `json:"input"      gencodec:"required"`
}
type ValTx struct {
	*valTxData
	from      atomic.Value
	hash      common.Hash
	time      time.Time
	Signature *bls.Sign `json:"sig" gencodec:"required"`
}

func (vt ValTx) Hash() common.Hash {
	return vt.hash
}

func (vt ValTx) SignTx(pri crypto.PrivateKey) error {
	prv, ok := pri.(*bls.SecretKey)
	if !ok {
		return fmt.Errorf("invalid micro transaction private key for singer")
	}
	bts, err := rlp.EncodeToBytes(vt.valTxData)
	if err != nil {
		return err
	}
	vt.Signature = prv.SignByte(bts)
	return nil
}
