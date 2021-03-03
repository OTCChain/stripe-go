package main

import (
	"context"
	"fmt"
	badger "github.com/ipfs/go-ds-badger"
	"github.com/ipfs/go-log/v2"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/discovery"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	ma "github.com/multiformats/go-multiaddr"
	"github.com/otcChain/chord-go/p2p"
)

type BootHost struct {
	host      host.Host
	dht       *dht.IpfsDHT
	discovery discovery.Discovery
	cancel    context.CancelFunc
	ctx       context.Context
}

func main() {

	log.SetAllLoggers(log.LevelInfo)
	conf := loadConfig()
	p2pPriKey := conf.loadKey()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	listenAddr, err := ma.NewMultiaddr(fmt.Sprintf("/ip4/0.0.0.0/tcp/%d", p2p.DefaultP2pPort))
	if err != nil {
		panic(err)
	}
	p2pHost, err := libp2p.New(ctx,
		libp2p.ListenAddrs(listenAddr),
		libp2p.Identity(p2pPriKey),
		libp2p.EnableNATService(),
		libp2p.ForceReachabilityPublic(),
	)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Host create ID[%s] %s\n", p2pHost.ID().Pretty(), p2pHost.Addrs())

	var bootPeers = make([]peer.AddrInfo, 0)
	for _, id := range conf.Boots {
		addr, err := ma.NewMultiaddr(id)
		if err != nil {
			panic(err)
		}
		peerInfo, err := peer.AddrInfoFromP2pAddr(addr)
		if err != nil {
			panic(err)
		}
		bootPeers = append(bootPeers, *peerInfo)
		fmt.Println("add new boot strap:", id)
	}
	ds, err := badger.NewDatastore("boot_dht_table", nil)
	if err != nil {
		panic(err)
	}
	kademliaDHT, err := dht.New(ctx, p2pHost, dht.Datastore(ds), dht.BootstrapPeers(bootPeers...))
	if err != nil {
		panic(err)
	}
	fmt.Println("Bootstrapping the DHT")
	if err := kademliaDHT.Bootstrap(ctx); err != nil {
		panic(err)
	}
	bootHost := &BootHost{
		host:   p2pHost,
		ctx:    ctx,
		cancel: cancel,
		dht:    kademliaDHT,
	}

	if conf.Chat {
		bootHost.chat()
	}
	fmt.Println("=======>>boot node setup")
	select {}
}
