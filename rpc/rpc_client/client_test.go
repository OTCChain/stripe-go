package chordclient

import (
	"context"
	"fmt"
	"github.com/herumi/bls-eth-go-binary/bls"
	"github.com/otcChain/chord-go/chord/transaction"
	"github.com/otcChain/chord-go/common"
	"math/big"
	"testing"
)

func TestValidAddress(t *testing.T) {
	client, err := Dial("http://127.0.0.1:6666")
	if err != nil {
		t.Fatal(err)
	}
	var privateKey bls.SecretKey
	err = privateKey.DeserializeHexStr("066c6b1a28955a9089670d1e1386484f7370ef7b4f725876e72d82438de06c9e")
	if err != nil {
		t.Fatal(err)
	}

	publicKey := privateKey.GetPublicKey()

	fromAddress := common.PubKeyToAddr(publicKey)
	fmt.Println(fromAddress)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("=====>account nonce is ", nonce)
	value := big.NewInt(1000000000000000000) // in wei (1 eth)
	price := big.NewInt(1000000000000000000) // in wei (1 eth)
	gas := 456
	toAddress, err := common.HexToAddress("fed1gy3afwa745c88dxsznaw82trul3r2p5vsrhmms")
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	tx := transaction.NewValTx(nonce, &toAddress, value, price, chainID, uint64(gas), nil)

	if err := tx.SignTx(&privateKey); err != nil {
		t.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), tx)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("transaction sent: %s", tx.Hash().Hex())
}
