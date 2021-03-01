package utils

import (
	"fmt"
	"github.com/herumi/bls-eth-go-binary/bls"
	"github.com/rs/zerolog"
	"io/ioutil"
	"net"
	"os"
	"path/filepath"
	"time"
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

func writeTemporaryKeyFile(file string, content []byte) (string, error) {
	// Create the keystore directory with appropriate permissions
	// in case it is not present yet.
	const dirPerm = 0700
	if err := os.MkdirAll(filepath.Dir(file), dirPerm); err != nil {
		return "", err
	}
	// Atomic write: create a temporary hidden file first
	// then move it into place. TempFile assigns mode 0600.
	f, err := ioutil.TempFile(filepath.Dir(file), "."+filepath.Base(file)+".tmp")
	if err != nil {
		return "", err
	}
	if _, err := f.Write(content); err != nil {
		_ = f.Close()
		_ = os.Remove(f.Name())
		return "", err
	}
	_ = f.Close()
	return f.Name(), nil
}

func WriteKeyFile(file string, content []byte) error {
	name, err := writeTemporaryKeyFile(file, content)
	if err != nil {
		return err
	}
	return os.Rename(name, file)
}

func ToISO8601(t time.Time) string {
	var tz string
	name, offset := t.Zone()
	if name == "UTC" {
		tz = "Z"
	} else {
		tz = fmt.Sprintf("%03d00", offset/3600)
	}
	return fmt.Sprintf("%04d-%02d-%02dT%02d-%02d-%02d.%09d%s",
		t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), tz)
}
