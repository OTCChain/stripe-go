package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/libp2p/go-libp2p"
	ma "github.com/multiformats/go-multiaddr"
	"github.com/otcChain/chord-go/common"
	"github.com/otcChain/chord-go/p2p"
	"github.com/otcChain/chord-go/wallet"
)

func main() {

	addr := flag.String("file", ".", "use -addr=[ADDRESS] the address of the key file")

	filePath := flag.String("file", ".", "use -file=[FILEPATH] to load the key file")
	password := flag.String("pw", "", "use -pw=[PWD]  [PWD] is the password for the key file")
	flag.Parse()

	if *password == "" {
		panic("invalid password")
	}

	ks := wallet.NewKeyStore("")
	path := ks.JoinPath(*filePath)
	fedAddr := common.HexToAddress(*addr)
	_, err := ks.GetKey(fedAddr, path, *password)
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
		//libp2p.Identity(key.PrivateKey),
		libp2p.EnableNATService(),
		libp2p.ForceReachabilityPublic(),
	)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello World, my second hosts ID is %s\n", p2pHost.ID())
}
