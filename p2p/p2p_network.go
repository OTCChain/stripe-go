package p2p

import (
	"github.com/herumi/bls-eth-go-binary/bls"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/host"
	libPS "github.com/libp2p/go-libp2p-pubsub"
	"sync"
)

////import "github.com/libp2p/go-libp2p"
//import "github.com/libp2p/go-libp2p-core/host"
//
////import "github.com/libp2p/go-libp2p-core/network"
////import "github.com/libp2p/go-libp2p-core/peer"
////import "github.com/libp2p/go-libp2p-core/peerstore"
//import libPS "github.com/libp2p/go-libp2p-pubsub"
//
////import "github.com/multiformats/go-multiaddr"
//import "github.com/herumi/bls-eth-go-binary/bls"

type NetworkV1 struct {
	host            host.Host
	pubSub          *libPS.PubSub
	priKey          crypto.PrivKey
	lock            sync.Mutex
	blockList       libPS.Blacklist
	ConsensusPubKey *bls.PublicKey
}

func (nt *NetworkV1) Setup(cfg *Config) error {
	return nil
}

func newNetwork() *NetworkV1 {
	n := &NetworkV1{}
	return n
}
