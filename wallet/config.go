package wallet

import "fmt"

const KeyStoreScheme = "keystore"
const TestKeyStoreScheme = "test_keystore"

type Config struct {
	Dir string `json:"keystore"`
}

func (c Config) String() string {
	s := fmt.Sprintf("\n<-------------P2p Config------------")
	s += fmt.Sprintf("\nkey store dir:%20s", c.Dir)
	s += fmt.Sprintf("\n----------------------------------->\n")
	return s
}

var instConf *Config = nil

func InitConfig(c *Config) {
	instConf = c
}
