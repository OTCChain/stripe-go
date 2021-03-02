package wallet

import (
	"crypto/ed25519"
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

func TestCastEdKey(t *testing.T) {
	k := NewKey()
	pri := k.PrivateKey.Serialize()
	var edPri = ed25519.NewKeyFromSeed(pri)
	t.Logf("edkey:%x\n", edPri)

	var edPriS = ed25519.NewKeyFromSeed(pri)
	t.Logf("edkey:%x\n", edPriS)

	msg := "hello world"
	sig := ed25519.Sign(edPri, []byte(msg))
	t.Logf("sig 1:%x\n", sig)
	edPub := edPri.Public().(ed25519.PublicKey)
	if !ed25519.Verify(edPub, []byte(msg), sig) {
		t.Fatal("convert 1 failed")
	}

	sig2 := ed25519.Sign(edPriS, []byte(msg))
	t.Logf("sig 2:%x\n", sig2)
	edPub2 := edPriS.Public().(ed25519.PublicKey)
	if !ed25519.Verify(edPub2, []byte(msg), sig2) {
		t.Fatal("convert 2 failed")
	}
	t.Logf("TestCastEdKey success=>:")
}

func TestCastP2PKey(t *testing.T) {
	k := NewKey()
	p2pKey, err := k.CastP2pKey()
	if err != nil {
		t.Fatal(err)
	}
	bts, err := p2pKey.Bytes()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("fisrt time get:[%x]\n", bts)
	for i := 0; i < 20; i++ {
		_, err = k.CastP2pKey()
		if err != nil {
			t.Fatal(err)
		}
	}
	p2pKey, err = k.CastP2pKey()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("22th get:[%x]\n", bts)
}
