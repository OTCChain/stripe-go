package main

import (
	"flag"
	"github.com/otcChain/chord-go/wallet"
)

func main() {
	dir := flag.String("dir", ".", "directory to save key file")
	password := flag.String("pw", "", "password to encrypt the key file")

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
