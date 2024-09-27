package p2p

import (
	"context"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-pubsub"
	"sync"
)

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
	MSNodeMsg   MessageChannel = "/0.1/Global/NODE"
	MSDebug     MessageChannel = "/0.1/Global/TEST"
)

var SystemTopics = []MessageChannel{MSConsensus, MSNodeMsg, MSDebug}

func consensusMsgValidator(ctx context.Context, peer peer.ID, msg *pubsub.Message) pubsub.ValidationResult {
	return pubsub.ValidationAccept
}

func nodeMsgValidator(ctx context.Context, peer peer.ID, msg *pubsub.Message) pubsub.ValidationResult {
	return pubsub.ValidationAccept
}
