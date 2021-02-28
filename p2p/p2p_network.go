package p2p

import (
	"context"
	"fmt"
	"github.com/herumi/bls-eth-go-binary/bls"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/host"
	libPS "github.com/libp2p/go-libp2p-pubsub"
	ma "github.com/multiformats/go-multiaddr"
	"sync"
)

type NetworkV1 struct {
	p2pHost         host.Host
	pubSub          *libPS.PubSub
	priKey          crypto.PrivKey
	lock            sync.Mutex
	blockList       libPS.Blacklist
	ConsensusPubKey *bls.PublicKey
	ctxCancel       context.CancelFunc
}

func newNetwork() *NetworkV1 {

	listenAddr, err := ma.NewMultiaddr(fmt.Sprintf("/tcp/%d", config.Port))
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	node, err := libp2p.New(ctx,
		libp2p.ListenAddrs(listenAddr),
		libp2p.Ping(false),
	)
	n := &NetworkV1{
		p2pHost:   node,
		ctxCancel: cancel,
	}
	return n
}

func (nt *NetworkV1) SetUp() error {
	return nil
}
func (nt *NetworkV1) Destroy() error {
	return nil
}
