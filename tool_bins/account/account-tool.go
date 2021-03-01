package main

import (
	"flag"
	"github.com/otcChain/chord-go/wallet"
)

func main() {
	dir := flag.String("dir", ".", "use -dir=[Directory] to save file to [Directory]")
	password := flag.String("pw", "", "use -pw=123")

	flag.Parse()

	if *password == "" {
		panic("password can't be empty")
	}

	key := wallet.NewKey()
	ks := wallet.NewKeyStore(*dir)
	if err := ks.StoreKey(key, *password); err != nil {
		panic(err)
	}
}
