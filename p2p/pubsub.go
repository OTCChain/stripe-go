package p2p

import (
	"context"
	coreDisc "github.com/libp2p/go-libp2p-core/discovery"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-discovery"
	"github.com/libp2p/go-libp2p-kad-dht"
	"github.com/libp2p/go-libp2p-pubsub"
	"github.com/otcChain/chord-go/utils"
	"sync"
)

type PubSub struct {
	ctx    context.Context
	lock   sync.RWMutex
	topics map[MessageChannel]*pubsub.Topic
	dht    *dht.IpfsDHT
	pubSub *pubsub.PubSub
	disc   coreDisc.Discovery
}

func newPubSub(ctx context.Context, h host.Host) (*PubSub, error) {
	dhtOpts, err := config.dhtOpts()
	kademliaDHT, err := dht.New(ctx, h, dhtOpts...)
	if err != nil {
		return nil, err
	}
	disc := discovery.NewRoutingDiscovery(kademliaDHT)
	psOption := config.pubSubOpts(disc)
	ps, err := pubsub.NewGossipSub(ctx, h, psOption...)
	if err != nil {
		return nil, err
	}

	topics := make(map[MessageChannel]*pubsub.Topic)
	for _, topID := range SystemTopics {
		topic, err := ps.Join(string(topID))
		if err != nil {
			return nil, err
		}
		topics[topID] = topic
	}

	return &PubSub{
		ctx:    ctx,
		dht:    kademliaDHT,
		pubSub: ps,
		disc:   disc,
		topics: topics,
	}, nil
}

func (s *PubSub) start() error {
	if err := s.dht.Bootstrap(s.ctx); err != nil {
		return err
	}

	for id, topic := range s.topics {
		sub, err := topic.Subscribe()
		if err != nil {
			return err
		}
		go s.readingMessage(id, sub)
	}
	return nil
}

func (s *PubSub) readingMessage(id MessageChannel, sub *pubsub.Subscription) {

	utils.LogInst().Info().Msgf("[pubSub] start reading [%s] message:", id)

	defer func() {
		s.lock.Lock()
		delete(s.topics, id)
		s.lock.Unlock()
		utils.LogInst().Warn().Msgf("remove topic [%s] from system", id)
	}()

	for {
		msg, err := sub.Next(s.ctx)
		if err != nil {
			utils.LogInst().Warn().Err(err)
			return
		}
		utils.LogInst().Debug().Msg(msg.String())
	}

}
