package wallet

import (
	"encoding/json"
	"fmt"
	"github.com/herumi/bls-eth-go-binary/bls"
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
func TestKey(t *testing.T) {
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
	key2, err := DecryptKey(bs, auth)
	if key.Address != key2.Address {
		t.Fatal("address is not same")
	}

	if !key.PrivateKey.IsEqual(key2.PrivateKey) {
		t.Fatal("private key is not same")
	}
	fmt.Println("case 2 success=>")
}
