package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/libp2p/go-libp2p-core/crypto"
	ma "github.com/multiformats/go-multiaddr"
	"github.com/otcChain/chord-go/wallet"
	"golang.org/x/crypto/ssh/terminal"
	"os"
)

type addrList []ma.Multiaddr

type BootConfig struct {
	Boots    []string `json:"boot"`
	KeyPath  string   `json:"keyPath"`
	Password string   `json:"password"`
	Chat     bool     `json:"test"`
}

const (
	DefaultConfigFileName = "boot.json"
)

//"/ip4/192.168.30.13/tcp/8888/p2p/12D3KooWRDCMDA11ypS2FM5ZxgG8QRMmuFTxExhcz9XixW2JMVSX",
//"/ip4/192.168.30.214/tcp/8888/p2p/12D3KooWBVTZ6qpuf2B5NqRrVxxDxUM7oPVWcdHa292SundjQpHH",
//"/ip4/202.182.101.145/tcp/8888/p2p/12D3KooWH1vt62wMAzSBHaAhH273MV8hnNuwF7jrDWptGzGFzPNe",
func initDefaultConfig() {
	bc := &BootConfig{
		Boots:    make([]string, 1),
		KeyPath:  "key.json",
		Password: "",
	}
	bc.Boots[0] = "/ip4/202.182.101.145/tcp/8888/p2p/12D3KooWH1vt62wMAzSBHaAhH273MV8hnNuwF7jrDWptGzGFzPNe"

	bts, err := json.MarshalIndent(bc, "", "\t")
	if err != nil {
		panic(err)
	}
	if err := os.WriteFile(DefaultConfigFileName, bts, 0644); err != nil {
		panic(err)
	}

	fmt.Println("Create default init config success!")
}

func loadConfig() *BootConfig {

	help := flag.Bool("h", false, "boot_node.[mac|lnx|arm|exe] -h")
	isInit := flag.Bool("init", false, "boot_node.[mac|lnx|arm|exe] -init")
	confFile := flag.String("c", DefaultConfigFileName, "use -c [Config File] to load the config file")
	keyPath := flag.String("key", "", "use -key=[FILEPATH] to load the key file")
	password := flag.String("p", "", "use -p=[PWD]  [PWD] is the password for the key file")
	chatTest := flag.Bool("chat", false, "use -chat to show test channel")
	flag.Parse()

	if *isInit {
		initDefaultConfig()
		os.Exit(0)
	}

	helpStr := "boot_node.[mac|lnx|arm|exe] -f ./wallet_key_file.json \n" +
		"-key the wallet key file of the boot node" +
		"-p Password of the wallet key\n" +
		"-c Config file of the boot node\n" +
		"-init create a default config file with name 'boot.json'\n" +
		"-h this description\n"
	if *help {
		fmt.Print(helpStr)
		os.Exit(0)
	}

	bts, err := os.ReadFile(*confFile)
	if err != nil {
		panic(err)
	}

	conf := &BootConfig{}
	if err := json.Unmarshal(bts, conf); err != nil {
		panic(err)
	}

	if *password != "" {
		conf.Password = *password
	}
	if conf.Password == "" {
		fmt.Println("Password=>")
		pw, err := terminal.ReadPassword(int(os.Stdin.Fd()))
		if err != nil {
			panic(err)
		}
		conf.Password = string(pw)
	}

	if *keyPath != "" {
		conf.KeyPath = *keyPath
	}

	if *chatTest {
		conf.Chat = true
	}

	fmt.Println(conf)
	return conf
}
func (conf *BootConfig) String() string {
	s := fmt.Sprintf("\n==============boot config==============")
	s += fmt.Sprintf("\nKeyPath  %20.s", conf.KeyPath)
	s += fmt.Sprintf("\nwith chat test%20.t", conf.Chat)
	s += fmt.Sprintf("\nno password   %20.t", conf.Password == "")
	s += fmt.Sprintf("\n|--------boot strap node peers(%d)\n|", len(conf.Boots))
	for i, boot := range conf.Boots {
		s += fmt.Sprintf("\n|node[%d]=%s", i, boot)
	}
	s += fmt.Sprintf("\n|\n|--------")
	s += fmt.Sprintf("\n=======================================\n")
	return s
}

func (conf *BootConfig) loadKey() crypto.PrivKey {
	ks := wallet.NewKeyStore("")
	path := ks.JoinPath(conf.KeyPath)
	walletKey, err := ks.GetRawKey(path, conf.Password)
	if err != nil {
		panic(err)
	}
	p2pPriKey, err := walletKey.CastP2pKey()
	if err != nil {
		panic(err)
	}
	return p2pPriKey
}
