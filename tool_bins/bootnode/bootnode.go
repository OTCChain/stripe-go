package main

import "flag"

func main() {

	//file := flag.String("file", ".", "directory to save key file")
	password := flag.String("pw", "", "password to encrypt the key file")

	flag.Parse()

	if *password == "" {
		panic("invalid password")
	}

}
