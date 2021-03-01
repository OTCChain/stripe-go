package wallet

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/herumi/bls-eth-go-binary/bls"
	"github.com/otcChain/chord-go/common"
	"testing"
)

func init() {
	if err := bls.Init(bls.BLS12_381); err != nil {
		panic(err)
	}
	if err := bls.SetETHmode(bls.EthModeDraft07); err != nil {
		panic(err)
	}
}
func TestKeyNew(t *testing.T) {
	k := NewKey()

	bs, err := json.MarshalIndent(k, "", "\t")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(bs))
	pub := k.PrivateKey.GetPublicKey()
	fmt.Printf("case 1 success=> pub:%x\n", pub.Serialize())
}

func TestKeyAuth(t *testing.T) {
	var auth = "123"
	key := NewKey()
	bs, err := key.Encrypt(auth)
	if err != nil {
		t.Fatal(err)
	}
	cipherTxt := string(bs)
	fmt.Println(cipherTxt)
	parsedKey, err := DecryptKey(bs, auth)
	if err != nil {
		t.Fatal(err)
	}
	if key.Address != parsedKey.Address {
		t.Fatal("address is not same")
	}

	if !key.PrivateKey.IsEqual(parsedKey.PrivateKey) {
		t.Fatal("private key is not same")
	}
	fmt.Println("case 2 success=>")
}

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
	path := ks.JoinPath("UTC--2021-03-01T07-50-07.030501000Z--fed15cpkgmn8nv56ja47h8qtwaq8tsyx8qv6pzsvwg")
	addr := common.HexToAddress("fed15cpkgmn8nv56ja47h8qtwaq8tsyx8qv6pzsvwg")
	key, err := ks.GetKey(addr, path, auth)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(key)
	if key.Address != addr {
		t.Fatal("load key failed")
	}
}
