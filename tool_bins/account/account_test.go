package main

import (
	"encoding/hex"
	"fmt"
	"github.com/herumi/bls-eth-go-binary/bls"
	"github.com/otcChain/chord-go/common"
	"github.com/otcChain/chord-go/wallet"
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

func TestAccount(t *testing.T) {
	pri := wallet.GenerateKey()
	pub := pri.GetPublicKey()
	addr := common.PubKeyToAddr(pub)
	fmt.Println(addr.Hex())
	fmt.Println(addr.String())
	if !(common.IsFedAddress(addr.String())) {
		t.Fatal("wallet string test failed")
	}
}

func TestCheckSign(t *testing.T) {
	fmt.Println("sample1")
	var sec bls.SecretKey
	sec.SetByCSPRNG()
	msg := []byte("abc")
	pub := sec.GetPublicKey()
	sig := sec.SignByte(msg)

	if !sig.VerifyByte(pub, msg) {
		t.Fatal("Signature should verify pass")
	}
	fmt.Println("sample1 success")
}

func TestCheckPub(t *testing.T) {
	var sec bls.SecretKey
	sec.SetByCSPRNG()
	fmt.Printf("sec:%s\n", sec.SerializeToHexStr())
	pub := sec.GetPublicKey()
	fmt.Printf("1.pub:%s\n", pub.SerializeToHexStr())
	fmt.Printf("1.pub x=%x\n", pub.Serialize())
	var P = bls.CastFromPublicKey(pub)
	bls.G1Normalize(P, P)
	fmt.Printf("2.pub:%s\n", pub.SerializeToHexStr())
	fmt.Printf("2.pub x=%x\n", pub.Serialize())
	fmt.Printf("P.X=%x\n", P.X.Serialize())
	fmt.Printf("P.Y=%x\n", P.Y.Serialize())
	fmt.Printf("P.Z=%x\n", P.Z.Serialize())
}

func TestLittleEndian(t *testing.T) {
	fmt.Printf("sample3\n")
	var sec bls.SecretKey
	b := make([]byte, 64)
	for i := 0; i < len(b); i++ {
		b[i] = 0xff
	}
	err := sec.SetLittleEndianMod(b)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("sec=%x\n", sec.Serialize())
}

func TestDecode(t *testing.T) {
	fmt.Printf("sample4\n")
	var sec bls.SecretKey
	//secByte, err := hex.DecodeString("4aac41b5cb665b93e031faa751944b1f14d77cb17322403cba8df1d6e4541a4d")//
	secByte, err := hex.DecodeString("28464487afefe15b5d92d3fa16cfee806ba078409ea5953fda52be82d060b3cb") //
	if err != nil {
		t.Fatal(err)
	}
	if err := sec.Deserialize(secByte); err != nil {
		t.Fatal(err)
	}
	msg := []byte("message to be signed.")
	fmt.Printf("sec:%x\n", sec.Serialize())
	pub := sec.GetPublicKey()
	fmt.Printf("pub:%x\n", pub.Serialize())
	sig := sec.SignByte(msg)
	fmt.Printf("sig:%x\n", sig.Serialize())
}

func TestLoadKey(t *testing.T) {
	var auth = "123"
	ks := wallet.NewLightKeyStore("", true)
	path := ks.JoinPath("UTC--2021-03-01T08-34-08.662288000Z--fed1djkthhrcql2f540fypqyp5df99s73dmwaafxkv")
	addr := common.HexToAddress("fed1djkthhrcql2f540fypqyp5df99s73dmwaafxkv")
	key, err := ks.GetKey(addr, path, auth)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(key)
	if key.Address != addr {
		t.Fatal("load key failed")
	}
	t.Log("=====>TestLoadKey success", key.Address.String())
}
