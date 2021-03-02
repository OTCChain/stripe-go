package node

import "fmt"

type Config struct {
}

func (c Config) String() string {
	s := fmt.Sprintf("\n<-------------Node Config-----------")
	s += fmt.Sprintf("\n----------------------------------->\n")
	return s
}

var config *Config = &Config{}

func InitConfig(c *Config) {
	config = c
}
func InitDefaultConfig() *Config {
	return &Config{}
}
