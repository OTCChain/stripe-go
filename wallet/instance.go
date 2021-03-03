package wallet

import (
	"fmt"
	"github.com/herumi/bls-eth-go-binary/bls"
	"github.com/otcChain/chord-go/common"
	"sync"
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
	Contains(key Key) bool
	Active(password, selectAddr string) error
	KeyInUsed() *Key
}

var _instance Wallet
var once sync.Once

func Inst() Wallet {
	once.Do(func() {
		_instance = newWallet()
	})
	return _instance
}

func newWallet() Wallet {
	cw := &ChordWalletV1{
		keys: make(map[common.Address]*Key),
	}
	return cw
}

type ChordWalletV1 struct {
	activeKey *Key
	sync.RWMutex
	keys map[common.Address]*Key
}

func (c *ChordWalletV1) Active(password, addr string) error {
	if instConf == nil {
		return fmt.Errorf("please init wallet instance config first")
	}

	ks := NewKeyStore(instConf.Dir)
	validKeyFiles, err := ks.ValidKeyFiles()
	if err != nil {
		return err
	}
	var selFile string
	if addr == "" && len(validKeyFiles) == 1 {
		addr = validKeyFiles[0]
	}

	address, _ := common.HexToAddress(addr)
	key, err := ks.GetKey(address, selFile, password)
	if err != nil {
		return err
	}

	c.activeKey = key
	c.Lock()
	c.keys[address] = key
	c.Unlock()
	return nil
}

func (c *ChordWalletV1) KeyInUsed() *Key {
	return c.activeKey
}

func (c *ChordWalletV1) Close() error {
	panic("implement me")
}

func (c *ChordWalletV1) Keys() []Key {
	panic("implement me")
}

func (c *ChordWalletV1) Contains(key Key) bool {
	panic("implement me")
}
