package utils

import (
	"fmt"
	"github.com/rs/zerolog"
	"os"
)

func FileExists(fileName string) bool {
	fileInfo, err := os.Lstat(fileName)
	if fileInfo != nil || (err != nil && !os.IsNotExist(err)) {
		return true
	}
	return false
}

type Config struct {
	LogLevel zerolog.Level
}

func (c Config) String() string {
	s := fmt.Sprintf("\n-------------P2p Config------------")
	s += fmt.Sprintf("\nlog level:%20s", c.LogLevel)
	s += fmt.Sprintf("\n-----------------------------------")
	return s
}

var config *Config = nil

func InitConfig(c *Config) {
	config = c
}
