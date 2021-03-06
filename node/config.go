package node

import (
	"fmt"
	"math/big"
)

var (
	MainChain = big.NewInt(1)
	TestChain = big.NewInt(2)
)

type Config struct {
	ChainID *big.Int
}

func (c Config) String() string {
	s := fmt.Sprintf("\n<-------------Node Config-----------")
	s += fmt.Sprintf("\n*chain id:			%s", c.ChainID.String())
	s += fmt.Sprintf("\n----------------------------------->\n")
	return s
}

var _nodeConfig *Config = &Config{}

func InitConfig(c *Config) {
	_nodeConfig = c
}
func DefaultConfig(isMain bool, base string) *Config {
	var chainID *big.Int
	if isMain {
		chainID = MainChain
	} else {
		chainID = TestChain
	}

	return &Config{
		ChainID: chainID,
	}
}
