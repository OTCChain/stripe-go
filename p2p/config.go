package p2p

import (
	"fmt"
)

const (
	DefaultP2pPort = 8888
)

type Config struct {
	Port int16 `json:"port"`
}

func (c Config) String() string {
	s := fmt.Sprintf("\n<-------------P2p Config------------")
	s += fmt.Sprintf("\nport:%20d", c.Port)
	s += fmt.Sprintf("\n----------------------------------->\n")
	return s
}

var config *Config = nil

func InitConfig(c *Config) {
	config = c
}
