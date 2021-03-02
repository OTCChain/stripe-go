package main

import (
	"bufio"
	"fmt"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/protocol"
	discovery "github.com/libp2p/go-libp2p-discovery"
	"os"
)

var (
	ProtocolID   = protocol.ID("/chord/boot")
	RendezvousID = "block_syncing"
	inputCh      = make([]chan string, 0)
)

func readData(rw *bufio.ReadWriter) {
	for {
		str, err := rw.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading from buffer", err)
			return
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

func writeData(rw *bufio.ReadWriter, ch <-chan string) {

	for {
		sendData := <-ch
		_, err := rw.WriteString(fmt.Sprintf("%s\n", sendData))
		if err != nil {
			fmt.Println("Error writing to buffer", err)
			return
		}
		err = rw.Flush()
		if err != nil {
			fmt.Println("Error flushing buffer", err)
		}
	}
}
func handleStream(stream network.Stream) {
	fmt.Println("Got a new stream!")

	// Create a buffer stream for non blocking read and write.
	rw := bufio.NewReadWriter(bufio.NewReader(stream), bufio.NewWriter(stream))
	in := make(chan string)
	go readData(rw)
	go writeData(rw, in)
	inputCh = append(inputCh, in)

	// 'stream' will stay open until you close it (or the other side closes it).
}
func (bh *BootHost) chat() {
	bh.host.SetStreamHandler(ProtocolID, handleStream)
	fmt.Println("Announcing ourselves...")
	disc := discovery.NewRoutingDiscovery(bh.dht)
	bh.discovery = disc
	duration, err := disc.Advertise(bh.ctx, RendezvousID)
	if err != nil {
		fmt.Println("advertise self err:", err)
	}
	fmt.Println("Announcing ourselves...", duration)
	fmt.Println("Searching for other peers...")

	peerChan, err := disc.FindPeers(bh.ctx, RendezvousID)
	if err != nil {
		panic(err)
	}

	for peerAddr := range peerChan {
		if peerAddr.ID == bh.host.ID() {
			fmt.Println("eee myself...")
			continue
		}
		fmt.Println("Found peer:", peerAddr)
		stream, err := bh.host.NewStream(bh.ctx, peerAddr.ID, ProtocolID)
		if err != nil {
			fmt.Println("Connection failed for peer:", peerAddr)
			continue
		}
		rw := bufio.NewReadWriter(bufio.NewReader(stream), bufio.NewWriter(stream))

		in := make(chan string)
		go writeData(rw, in)
		go readData(rw)
		inputCh = append(inputCh, in)

		fmt.Println("Connected to:", peerAddr)
	}
	fmt.Println("=======>>boot node setup")

	stdReader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		sendData, err := stdReader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading from stdin", err)
			continue
		}
		for _, ch := range inputCh {
			ch <- sendData
		}
	}
}
