package p2p

import (
	"fmt"
	badger "github.com/ipfs/go-ds-badger"
	"github.com/ipfs/go-log/v2"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/discovery"
	"github.com/libp2p/go-libp2p-core/peer"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	"github.com/libp2p/go-libp2p-pubsub"
	ma "github.com/multiformats/go-multiaddr"
	"github.com/otcChain/chord-go/utils"
	"github.com/otcChain/chord-go/wallet"
	"path/filepath"
	"runtime"
)

const (
	DefaultP2pPort           = 8888
	DefaultMaxMessageSize    = 1 << 21
	DefaultOutboundQueueSize = 64
	DefaultValidateQueueSize = 512

	DefaultConsensusTopicThreadSize = 1 << 13
	DefaultOtherTopicThreadSize     = 1 << 11

	DHTPrefix = "chord"
)

var (
	MainP2pBoots = []string{"/ip4/202.182.101.145/tcp/8888/p2p/12D3KooWH1vt62wMAzSBHaAhH273MV8hnNuwF7jrDWptGzGFzPNe"}
	TestP2pBoots = []string{"/ip4/202.182.101.145/tcp/8888/p2p/12D3KooWH1vt62wMAzSBHaAhH273MV8hnNuwF7jrDWptGzGFzPNe",
		"/ip4/192.168.30.214/tcp/8888/p2p/12D3KooWBVTZ6qpuf2B5NqRrVxxDxUM7oPVWcdHa292SundjQpHH"}
)

type pubSubConfig struct {
	MaxMsgSize          int `json:"max_msg_size"`
	MaxValidateQueue    int `json:"validate_queue_size"`
	MaxOutQueue         int `json:"out_queue_size"`
	MaxConsTopicThread  int `json:"consensus_topic_threads"`
	MaxOtherTopicThread int `json:"other_topic_threads"`
}

func (c *pubSubConfig) String() string {
	s := fmt.Sprintf("\n<*******pub sub*********")
	s += fmt.Sprintf("\n*max message:			%d", c.MaxMsgSize)
	s += fmt.Sprintf("\n*max validate queue size:	%d", c.MaxValidateQueue)
	s += fmt.Sprintf("\n*max out queue size:		%d", c.MaxOutQueue)
	s += fmt.Sprintf("\n*max consensus topic thread:	%d", c.MaxConsTopicThread)
	s += fmt.Sprintf("\n*max common topic thread:	%d", c.MaxOtherTopicThread)
	s += fmt.Sprintf("\n*************************\n")
	return s
}

type dhtConfig struct {
	DataStoreFile string   `json:"cache_dir"`
	Boots         []string `json:"bootstrap"`
}

func (c *dhtConfig) String() string {
	s := fmt.Sprintf("\n<**********dht***********")
	s += fmt.Sprintf("\n*dht cache dir:%s", c.DataStoreFile)
	s += fmt.Sprintf("\n*boot strap nodes:%d", len(c.Boots))
	for _, boot := range c.Boots {
		s += fmt.Sprintf("\n%s", boot)
	}
	s += fmt.Sprintf("\n*************************\n")
	return s
}

type Config struct {
	Port     int16         `json:"port"`
	LogLevel log.LogLevel  `json:"log_level"`
	PsConf   *pubSubConfig `json:"pub_sub"`
	DHTConf  *dhtConfig    `json:"dht"`
}

func (c Config) String() string {
	s := fmt.Sprintf("\n<-------------P2p Config------------")
	s += fmt.Sprintf("\nport:		%d", c.Port)
	s += fmt.Sprintf("\nloglevl:	%d", c.LogLevel)
	s += fmt.Sprintf(c.PsConf.String())
	s += fmt.Sprintf(c.DHTConf.String())
	s += fmt.Sprintf("\n----------------------------------->\n")
	return s
}

var config *Config = nil

func DefaultConfig(isMain bool, base string) *Config {
	var (
		level  log.LogLevel
		boots  []string
		dhtDir string
	)
	if isMain {
		boots = MainP2pBoots
		level = log.LevelWarn
		dhtDir = filepath.Join(base, string(filepath.Separator), "dht_cache")
	} else {
		boots = TestP2pBoots
		level = log.LevelDebug
		dhtDir = filepath.Join(base, string(filepath.Separator), "dht_cache_test")

	}

	return &Config{
		Port:     DefaultP2pPort,
		LogLevel: level,
		PsConf: &pubSubConfig{
			MaxMsgSize:          DefaultMaxMessageSize,
			MaxValidateQueue:    DefaultValidateQueueSize,
			MaxOutQueue:         DefaultOutboundQueueSize,
			MaxConsTopicThread:  DefaultConsensusTopicThreadSize,
			MaxOtherTopicThread: DefaultOtherTopicThreadSize,
		},
		DHTConf: &dhtConfig{
			DataStoreFile: dhtDir,
			Boots:         boots,
		},
	}
}

func InitConfig(c *Config) {
	config = c
}

func (c *Config) initOptions() []libp2p.Option {
	listenAddr, err := ma.NewMultiaddr(fmt.Sprintf("/ip4/0.0.0.0/tcp/%d", config.Port))
	if err != nil {
		panic(err)
	}

	activeKey := wallet.Inst().KeyInUsed()
	if activeKey == nil {
		panic("no valid key right now")
	}
	key, err := activeKey.CastP2pKey()
	if err != nil {
		panic(err)
	}

	return []libp2p.Option{
		libp2p.ListenAddrs(listenAddr),
		libp2p.Identity(key),
		libp2p.EnableNATService(),
		libp2p.ForceReachabilityPublic(),
	}
}

func (c *Config) pubSubOpts(disc discovery.Discovery) []pubsub.Option {
	return []pubsub.Option{
		pubsub.WithValidateQueueSize(c.PsConf.MaxValidateQueue),
		pubsub.WithPeerOutboundQueueSize(c.PsConf.MaxOutQueue),
		pubsub.WithValidateWorkers(runtime.NumCPU() * 2),
		pubsub.WithValidateThrottle(c.PsConf.MaxConsTopicThread + c.PsConf.MaxOtherTopicThread),
		pubsub.WithMaxMessageSize(c.PsConf.MaxMsgSize),
		pubsub.WithDiscovery(disc),
	}
}

func (c *Config) dhtOpts() ([]dht.Option, error) {
	ds, err := badger.NewDatastore(c.DHTConf.DataStoreFile, nil)
	if err != nil {
		return nil, fmt.Errorf("cannot open Badger data store at %s, err:%s",
			c.DHTConf.DataStoreFile, err)
	}
	peers := make([]peer.AddrInfo, 0)

	for _, id := range c.DHTConf.Boots {
		addr, err := ma.NewMultiaddr(id)
		if err != nil {
			utils.LogInst().Warn().Str("invalid boot id", id)
			continue
		}
		peerInfo, err := peer.AddrInfoFromP2pAddr(addr)
		if err != nil {
			utils.LogInst().Warn().Str("parse failed for boot id", id)
			continue
		}
		peers = append(peers, *peerInfo)
	}
	if len(peers) == 0 {
		return nil, fmt.Errorf("no invalid bootstrap node")
	}

	return []dht.Option{
		dht.Datastore(ds),
		dht.ProtocolPrefix(DHTPrefix),
		dht.BootstrapPeers(peers...),
	}, nil
}
