package consensus

import "fmt"

type Config struct {
}

func (c Config) String() string {
	s := fmt.Sprintf("\n----------Consensus Config---------")
	s += fmt.Sprintf("\n-----------------------------------")
	return s
}

var config *Config = nil

func InitConfig(c *Config) {
	config = c
}
