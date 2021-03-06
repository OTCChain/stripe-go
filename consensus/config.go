package consensus

import "fmt"

type Config struct {
}

func (c Config) String() string {
	s := fmt.Sprintf("\n<----------Consensus Config---------")
	s += fmt.Sprintf("\n----------------------------------->\n")
	return s
}

var _consConfig *Config = &Config{}

func InitConfig(c *Config) {
	_consConfig = c
}
func DefaultConfig(isMain bool, baseDir string) *Config {
	return &Config{}
}
