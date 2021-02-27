package account

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/btcsuite/btcutil/bech32"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
)

const (
	ConfFileName     = "account.json"
	Bech32AddressHRP = "fed"
)

func IsFedAddress(s string) bool {
	hrp, bytes, err := decodeAndConvert(s)
	if err != nil || (hrp != Bech32AddressHRP) || len(bytes) != common.AddressLength {
		return false
	}
	return true
}
func convertAndEncode(hrp string, data []byte) (string, error) {
	converted, err := bech32.ConvertBits(data, 8, 5, true)
	if err != nil {
		return "", errors.Wrap(err, "encoding bech32 failed")
	}
	return bech32.Encode(hrp, converted)

}
func decodeAndConvert(bech string) (string, []byte, error) {
	hrp, data, err := bech32.Decode(bech)
	if err != nil {
		return "", nil, errors.Wrap(err, "decoding bech32 failed")
	}
	converted, err := bech32.ConvertBits(data, 5, 8, false)
	if err != nil {
		return "", nil, errors.Wrap(err, "decoding bech32 failed")
	}
	return hrp, converted, nil
}

func BuildAddrByETHAddr(hrp string, addr common.Address) (string, error) {
	return convertAndEncode(hrp, addr.Bytes())
}

func MustBuildByETHAddr(hrp string, addr common.Address) string {
	b32, err := BuildAddrByETHAddr(hrp, addr)
	if err != nil {
		panic(err)
	}
	return b32
}

func ToETHAddress(b32 string) (addr common.Address, err error) {
	hrp, bytes, err := decodeAndConvert(b32)
	if err != nil {
		err = fmt.Errorf("cannot decode %#v as bech32 address:%s", b32, err)
		return
	}
	if hrp != Bech32AddressHRP {
		err = fmt.Errorf("%#v is not a %#v address", b32, Bech32AddressHRP)
		return
	}
	if len(bytes) != common.AddressLength {
		err = fmt.Errorf("decoded bech32 %#v has invalid length %d",
			b32, len(bytes))
		return
	}
	addr.SetBytes(bytes)
	return
}

func MustToETHAddr(b32 string) common.Address {
	addr, err := ToETHAddress(b32)
	if err != nil {
		panic(err)
	}
	return addr
}

func FromETHAddress(addr common.Address) (string, error) {
	return BuildAddrByETHAddr(Bech32AddressHRP, addr)
}

func MustFromETHAddress(addr common.Address) string {
	b32, err := BuildAddrByETHAddr(Bech32AddressHRP, addr)
	if err != nil {
		panic(err)
	}
	return b32
}

func ParseAddr(s string) common.Address {
	if addr, err := ToETHAddress(s); err == nil {
		return addr
	}
	return common.HexToAddress(s)
}

func MustGeneratePrivateKey() *ecdsa.PrivateKey {
	key, err := crypto.GenerateKey()
	if err != nil {
		panic(err)
	}
	return key
}
