package rpc

import (
	"fmt"
	"time"
)

const (
	DefaultHttpPort     = 6666
	DefaultWsPort       = 6646
	DefaultReadTimeout  = 30 * time.Second
	DefaultWriteTimeout = 30 * time.Second
	DefaultIdleTimeout  = 120 * time.Second
	DefaultHost         = "localhost"
)

type Config struct {
	HttpEnabled  bool          `json:"http.en"`
	HttpIP       string        `json:"http.ip"`
	HttpPort     int16         `json:"http.port"`
	ReadTimeout  time.Duration `json:"http.r.timeout"`
	WriteTimeout time.Duration `json:"http.w.timeout"`
	IdleTimeout  time.Duration `json:"http.i.timeout"`

	WsEnabled bool   `json:"ws.en"`
	WsIP      string `json:"ws.ip"`
	WsPort    int16  `json:"ws.port"`
}

func (c Config) String() string {
	s := fmt.Sprintf("\n<-------------rpc Config------------")
	s += fmt.Sprintf("\nhttp enabled:%20t", c.HttpEnabled)
	s += fmt.Sprintf("\nhttp ip:%20s", c.HttpIP)
	s += fmt.Sprintf("\nhttp port:%20d", c.HttpPort)
	s += fmt.Sprintf("\nhttp read timeout:%20d", c.ReadTimeout)
	s += fmt.Sprintf("\nhttp writ timeout:%20d", c.WriteTimeout)
	s += fmt.Sprintf("\nhttp idle timeout:%20d", c.IdleTimeout)
	s += fmt.Sprintf("\nws enabled:%20t", c.WsEnabled)
	s += fmt.Sprintf("\nws ip:%20s", c.WsIP)
	s += fmt.Sprintf("\nws port:%20d", c.WsPort)
	s += fmt.Sprintf("\n----------------------------------->\n")
	return s
}

var _rpcConfig *Config = nil

func InitConfig(c *Config) {
	_rpcConfig = c
}

func DefaultConfig() *Config {

	return &Config{
		HttpEnabled:  false,
		HttpIP:       DefaultHost,
		HttpPort:     DefaultHttpPort,
		ReadTimeout:  DefaultReadTimeout,
		WriteTimeout: DefaultWriteTimeout,
		IdleTimeout:  DefaultIdleTimeout,

		WsEnabled: false,
		WsIP:      DefaultHost,
		WsPort:    DefaultWsPort,
	}
}
