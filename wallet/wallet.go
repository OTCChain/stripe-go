package wallet

import (
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/otcChain/chord-go/common"
	"math/big"
)

type Wallet interface {
	URL() string
	Open(passphrase string) error
	Close() error
	Keys() []Key
	Contains(key Key) bool
	Derive(path accounts.DerivationPath, pin bool) (Key, error)
	SignData(account Key, mimeType string, data []byte) ([]byte, error)
	SignDataWithPassphrase(key Key, passphrase, mimeType string, data []byte) ([]byte, error)
	SignText(key Key, text []byte) ([]byte, error)
	SignTextWithPassphrase(account Key, passphrase string, hash []byte) ([]byte, error)
	SignTx(key Key, tx *common.Transaction, chainID *big.Int) (*common.Transaction, error)
	SignTxWithPassphrase(account Key, passphrase string, tx *common.Transaction, chainID *big.Int) (*common.Transaction, error)
}
