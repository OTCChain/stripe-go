package wallet

import (
	"encoding/hex"
	"fmt"
	"github.com/otcChain/chord-go/common"
	"testing"
)

func TestKeyPath(t *testing.T) {
	var auth = "123"
	key := NewKey()
	fmt.Println(key.Address)
	fmt.Println(hex.EncodeToString(key.Address[:]))
	ks := NewLightKeyStore("key_dir", key.Light)
	if err := ks.StoreKey(key, auth); err != nil {
		t.Fatal(err)
	}
	fmt.Println("case 3 success=>")
}

func TestLoadKey(t *testing.T) {
	var auth = "123"
	ks := NewLightKeyStore("key_dir", false)
	path := "UTC--2021-03-01T07-50-07.030501000Z--fed15cpkgmn8nv56ja47h8qtwaq8tsyx8qv6pzsvwg"
	addr, err := common.HexToAddress("fed15cpkgmn8nv56ja47h8qtwaq8tsyx8qv6pzsvwg")
	if err != nil {
		t.Fatal(err)
	}
	key, err := ks.GetKey(addr, path, auth)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(key)
	if key.Address != addr {
		t.Fatal("load key failed")
	}
}
