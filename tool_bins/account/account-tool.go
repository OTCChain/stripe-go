package main

import (
	"flag"
	"fmt"
	"github.com/otcChain/chord-go/wallet"
	"os"
)

func main() {
	help := flag.Bool("h", false, "chord_acc.[mac|lnx|arm|exe] -h")
	dir := flag.String("d", ".", "use -d=[Directory] to save file to [Directory]")
	password := flag.String("p", "", "use -p=[PASSWORD]")

	flag.Parse()
	helpStr := "chord_acc.[mac|lnx|arm|exe] -d [Directory to save account] -pw [Password of the wallet key]"
	if *help {
		fmt.Println(helpStr)
		os.Exit(0)
	}

	if *password == "" {
		panic(helpStr)
	}

	key := wallet.NewKey()
	ks := wallet.NewKeyStore(*dir)
	if err := ks.StoreKey(key, *password); err != nil {
		panic(err)
	}
}
