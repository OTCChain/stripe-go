package rpc

const (
	DefaultCmdPort  = 8848
	DefaultHttpPort = 6666
	DefaultWsPort   = 6646
	DefaultHost     = "localhost"
	SrvNameCmd      = "cmd"
	SrvNameHttp     = "http"
	SrvNameWs       = "ws"
)

type Config struct {
	CmdEnabled bool   `json:"cmd.en"`
	CmdIP      string `json:"cmd.ip"`
	CmdPort    int16  `json:"cmd.port"`

	HttpEnabled bool   `json:"http.en"`
	HttpIP      string `json:"http.ip"`
	HttpPort    int16  `json:"http.port"`

	WsEnabled bool   `json:"ws.en"`
	WsIP      string `json:"ws.ip"`
	WsPort    int16  `json:"ws.port"`
}

var config *Config = nil

func InitConfig(c *Config) {
	config = c
}

func DefaultConfig() *Config {

	return &Config{
		CmdEnabled: true,
		CmdIP:      DefaultHost,
		CmdPort:    DefaultCmdPort,

		HttpEnabled: false,
		HttpIP:      DefaultHost,
		HttpPort:    DefaultHttpPort,

		WsEnabled: false,
		WsIP:      DefaultHost,
		WsPort:    DefaultWsPort,
	}
}
