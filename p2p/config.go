package p2p

import "fmt"

type Config struct {
}

func (c Config) String() string {
	s := fmt.Sprintf("\n-------------P2p Config------------")
	s += fmt.Sprintf("\n-----------------------------------")
	return s
}

var config *Config = nil

func InitConfig(c *Config) {
	config = c
}
