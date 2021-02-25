package p2p

import "sync"
//import "github.com/libp2p/go-libp2p"
import "github.com/libp2p/go-libp2p-core/crypto"
import "github.com/libp2p/go-libp2p-core/host"
//import "github.com/libp2p/go-libp2p-core/network"
//import "github.com/libp2p/go-libp2p-core/peer"
//import "github.com/libp2p/go-libp2p-core/peerstore"
import libPS "github.com/libp2p/go-libp2p-pubsub"
//import "github.com/multiformats/go-multiaddr"

type NetworkV1 struct {
	host      host.Host
	pubSub    *libPS.PubSub
	priKey    crypto.PrivKey
	lock      sync.Mutex
	blockList libPS.Blacklist
}

func (nt *NetworkV1) Setup(cfg *Config) error{
	return nil
}

func newNetwork() *NetworkV1{
	n := &NetworkV1{

	}
	return n
}