package wallet

import (
	"github.com/herumi/bls-eth-go-binary/bls"
	"github.com/otcChain/chord-go/common"
	"math/big"
)

func init() {
	if err := bls.Init(bls.BLS12_381); err != nil {
		panic(err)
	}
	if err := bls.SetETHmode(bls.EthModeDraft07); err != nil {
		panic(err)
	}
}

type Wallet interface {
	URL() string
	Open(passphrase string) error
	Close() error
	Keys() []Key
	Contains(key Key) bool
	SignData(account Key, mimeType string, data []byte) ([]byte, error)
	SignDataWithPassphrase(key Key, passphrase, mimeType string, data []byte) ([]byte, error)
	SignText(key Key, text []byte) ([]byte, error)
	SignTextWithPassphrase(account Key, passphrase string, hash []byte) ([]byte, error)
	SignTx(key Key, tx *common.Transaction, chainID *big.Int) (*common.Transaction, error)
	SignTxWithPassphrase(account Key, passphrase string, tx *common.Transaction, chainID *big.Int) (*common.Transaction, error)
}
