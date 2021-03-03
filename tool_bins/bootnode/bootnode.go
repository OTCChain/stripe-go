package main

import (
	"context"
	"fmt"
	"github.com/ipfs/go-log/v2"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/discovery"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	ma "github.com/multiformats/go-multiaddr"
	"github.com/otcChain/chord-go/p2p"
	"sync"
)

type BootHost struct {
	host      host.Host
	dht       *dht.IpfsDHT
	discovery discovery.Discovery
	cancel    context.CancelFunc
	ctx       context.Context
}

func (bh *BootHost) bootStrap(conf *BootConfig) {
	var bootPeers addrList
	for _, id := range conf.Boots {
		addr, err := ma.NewMultiaddr(id)
		if err != nil {
			panic(err)
		}
		bootPeers = append(bootPeers, addr)
	}

	var wg sync.WaitGroup
	for _, peerAddr := range bootPeers {
		peerInfo, _ := peer.AddrInfoFromP2pAddr(peerAddr)
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := bh.host.Connect(bh.ctx, *peerInfo); err != nil {
				fmt.Println("Connect to boot strap err:", err, peerInfo)
			} else {
				fmt.Println("Connection established with bootstrap node:", *peerInfo)
			}
		}()
	}
	wg.Wait()
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
	kademliaDHT, err := dht.New(ctx, p2pHost)
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
	//bootHost.bootStrap(conf)
	//fmt.Println("boot strap start up......")

	if conf.Chat {
		bootHost.chat()
	}
	fmt.Println("=======>>boot node setup")
	select {}
}
