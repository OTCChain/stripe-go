package common

import (
	"math/big"
	"sync/atomic"
)

type Transaction struct {
	data txdata
	hash atomic.Value
	size atomic.Value
	from atomic.Value
}

type txdata struct {
	AccountNonce uint64   `json:"nonce"      gencodec:"required"`
	Price        *big.Int `json:"gasPrice"   gencodec:"required"`
	GasLimit     uint64   `json:"gas"        gencodec:"required"`
	Payload      []byte   `json:"input"      gencodec:"required"`
}
