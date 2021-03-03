package p2p

import (
	"context"
	"fmt"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	libPS "github.com/libp2p/go-libp2p-pubsub"
	ma "github.com/multiformats/go-multiaddr"
	"github.com/otcChain/chord-go/utils"
	"github.com/otcChain/chord-go/wallet"
	"sync"
)

type NetworkV1 struct {
	p2pHost   host.Host
	pubSub    *libPS.PubSub
	lock      sync.Mutex
	blockList libPS.Blacklist
	ctxCancel context.CancelFunc
}

func newNetwork() *NetworkV1 {

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

	ctx, cancel := context.WithCancel(context.Background())
	node, err := libp2p.New(ctx,
		libp2p.ListenAddrs(listenAddr),
		libp2p.Identity(key),
		libp2p.EnableNATService(),
		libp2p.ForceReachabilityPublic(),
	)

	if err != nil {
		panic(err)
	}
	n := &NetworkV1{
		p2pHost:   node,
		ctxCancel: cancel,
	}
	peerInfo := peer.AddrInfo{
		ID:    node.ID(),
		Addrs: node.Addrs(),
	}
	addrs, err := peer.AddrInfoToP2pAddrs(&peerInfo)
	if err != nil {
		panic(err)
	}
	utils.LogInst().Print("======>libp2p node address:", addrs[0])
	return n
}

func (nt *NetworkV1) SetUp() error {
	return nil
}

func (nt *NetworkV1) Destroy() error {
	return nil
}
