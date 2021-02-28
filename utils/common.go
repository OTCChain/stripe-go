package utils

import (
	"fmt"
	"github.com/herumi/bls-eth-go-binary/bls"
	"github.com/rs/zerolog"
	"net"
	"os"
)

func init() {
	_ = bls.Init(bls.BLS12_381)
	for _, cidr := range []string{
		"127.0.0.0/8",    // IPv4 loopback
		"10.0.0.0/8",     // RFC1918
		"172.16.0.0/12",  // RFC1918
		"192.168.0.0/16", // RFC1918
		"::1/128",        // IPv6 loopback
		"fe80::/10",      // IPv6 link-local
	} {
		_, block, _ := net.ParseCIDR(cidr)
		privateNets = append(privateNets, block)
	}
}

func FileExists(fileName string) bool {
	fileInfo, err := os.Lstat(fileName)
	if fileInfo != nil || (err != nil && !os.IsNotExist(err)) {
		return true
	}
	return false
}

var privateNets []*net.IPNet

func IsPrivateIP(ip net.IP) bool {
	for _, block := range privateNets {
		if block.Contains(ip) {
			return true
		}
	}
	return false
}

type Config struct {
	LogLevel zerolog.Level
}

func (c Config) String() string {
	s := fmt.Sprintf("\n<-------------utils Config------------")
	s += fmt.Sprintf("\nlog level:%20s", c.LogLevel)
	s += fmt.Sprintf("\n----------------------------------->\n")
	return s
}

var config *Config = nil

func InitConfig(c *Config) {
	config = c
}
