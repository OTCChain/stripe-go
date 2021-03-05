package rpc

import "time"

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
