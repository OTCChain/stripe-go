package p2p

import "sync"

type ChordNetwork interface {
	LaunchUp() error
	Destroy() error
}

var _instance ChordNetwork
var once sync.Once

func Inst() ChordNetwork {
	once.Do(func() {
		_instance = newNetwork()
	})
	return _instance
}

type MessageChannel string

func (mc MessageChannel) String() string {
	return string(mc)
}

const (
	MSConsensus MessageChannel = "/0.1/Global/CONSENSUS"
	MSClientRpc MessageChannel = "/0.1/Global/CLIENT"
	MSDebug     MessageChannel = "/0.1/Global/TEST"
)

var SystemTopics = []MessageChannel{MSConsensus, MSClientRpc, MSDebug}
