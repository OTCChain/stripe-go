package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	dis "github.com/libp2p/go-libp2p-discovery"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	ma "github.com/multiformats/go-multiaddr"
	"github.com/otcChain/chord-go/p2p"
	"github.com/otcChain/chord-go/wallet"
	"os"
	"sync"
)

type addrList []ma.Multiaddr

var (
	ProtocolID     = protocol.ID("/chord/boot")
	BootstrapPeers addrList
	RendezvousID   = "block_syncing"
)

func readData(rw *bufio.ReadWriter) {
	for {
		str, err := rw.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading from buffer")
			panic(err)
		}

		if str == "" {
			return
		}
		if str != "\n" {
			// Green console colour: 	\x1b[32m
			// Reset console colour: 	\x1b[0m
			fmt.Printf("\x1b[32m%s\x1b[0m> ", str)
		}

	}
}

func writeData(rw *bufio.ReadWriter) {
	stdReader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		sendData, err := stdReader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading from stdin")
			panic(err)
		}

		_, err = rw.WriteString(fmt.Sprintf("%s\n", sendData))
		if err != nil {
			fmt.Println("Error writing to buffer")
			panic(err)
		}
		err = rw.Flush()
		if err != nil {
			fmt.Println("Error flushing buffer")
			panic(err)
		}
	}
}
func handleStream(stream network.Stream) {
	fmt.Println("Got a new stream!")

	// Create a buffer stream for non blocking read and write.
	rw := bufio.NewReadWriter(bufio.NewReader(stream), bufio.NewWriter(stream))

	go readData(rw)
	go writeData(rw)

	// 'stream' will stay open until you close it (or the other side closes it).
}

func main() {

	help := flag.Bool("h", false, "boot_node.[mac|lnx|arm|exe] -h")
	filePath := flag.String("f", ".", "use -f=[FILEPATH] to load the key file")
	password := flag.String("p", "", "use -p=[PWD]  [PWD] is the password for the key file")
	flag.Parse()

	helpStr := "boot_node.[mac|lnx|arm|exe] -f ./wallet_key_file.json -p [Password of the wallet key]"
	if *help {
		fmt.Println(helpStr)
		os.Exit(0)
	}

	for _, id := range []string{
		//"/ip4/192.168.30.13/tcp/8888/p2p/12D3KooWRDCMDA11ypS2FM5ZxgG8QRMmuFTxExhcz9XixW2JMVSX",
		//"/ip4/192.168.30.214/tcp/8888/p2p/12D3KooWBVTZ6qpuf2B5NqRrVxxDxUM7oPVWcdHa292SundjQpHH",
		"/ip4/202.182.101.145/tcp/8888/p2p/12D3KooWH1vt62wMAzSBHaAhH273MV8hnNuwF7jrDWptGzGFzPNe",
	} {
		addr, err := ma.NewMultiaddr(id)
		if err != nil {
			panic(err)
		}
		BootstrapPeers = append(BootstrapPeers, addr)
	}

	if *password == "" {
		panic(helpStr)
	}

	ks := wallet.NewKeyStore("")
	path := ks.JoinPath(*filePath)
	walletKey, err := ks.GetRawKey(path, *password)
	if err != nil {
		panic(err)
	}
	p2pPriKey, err := walletKey.CastP2pKey()
	if err != nil {
		panic(err)
	}

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

	p2pHost.SetStreamHandler(ProtocolID, handleStream)

	kademliaDHT, err := dht.New(ctx, p2pHost)
	if err != nil {
		panic(err)
	}

	fmt.Println("Bootstrapping the DHT")
	if err := kademliaDHT.Bootstrap(ctx); err != nil {
		panic(err)
	}

	var wg sync.WaitGroup
	for _, peerAddr := range BootstrapPeers {
		peerInfo, _ := peer.AddrInfoFromP2pAddr(peerAddr)
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := p2pHost.Connect(ctx, *peerInfo); err != nil {
				fmt.Println("Connect to boot strap err:", err)
			} else {
				fmt.Println("Connection established with bootstrap node:", *peerInfo)
			}
		}()
	}
	wg.Wait()

	fmt.Println("Announcing ourselves...")
	discovery := dis.NewRoutingDiscovery(kademliaDHT)
	duration, err := discovery.Advertise(ctx, RendezvousID)
	if err != nil {
		fmt.Println("advertise self err:", err)
	}
	fmt.Println("Announcing ourselves...", duration)
	fmt.Println("Searching for other peers...")

	peerChan, err := discovery.FindPeers(ctx, RendezvousID)
	if err != nil {
		panic(err)
	}
	for peerAddr := range peerChan {
		if peerAddr.ID == p2pHost.ID() {
			fmt.Println("eee myself...")
			continue
		}
		fmt.Println("Found peer:", peerAddr)
		stream, err := p2pHost.NewStream(ctx, peerAddr.ID, ProtocolID)
		if err != nil {
			fmt.Println("Connection failed for peer:", peerAddr)
			continue
		}
		rw := bufio.NewReadWriter(bufio.NewReader(stream), bufio.NewWriter(stream))

		go writeData(rw)
		go readData(rw)

		fmt.Println("Connected to:", peerAddr)
	}
	fmt.Println("=======>>boot node setup")
	select {}
}
