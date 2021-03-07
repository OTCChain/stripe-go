package common

import (
	"database/sql/driver"
	"fmt"
	"github.com/btcsuite/btcutil/bech32"
	"github.com/herumi/bls-eth-go-binary/bls"
	"reflect"
)

const (
	ChordAddressHRP = "fed"
	AddressLength   = 20
	HashLength      = 32
)

type Address [AddressLength]byte

var (
	InvalidAddr Address
	AddressT    = reflect.TypeOf(Address{})
)

func (a *Address) SetBytes(b []byte) {
	if len(b) > len(a) {
		b = b[len(b)-AddressLength:]
	}
	copy(a[AddressLength-len(b):], b)
}

func convertAndEncode(hrp string, data []byte) (string, error) {
	converted, err := bech32.ConvertBits(data, 8, 5, true)
	if err != nil {
		return "", fmt.Errorf("convertBits failed:%s", err)
	}
	return bech32.Encode(hrp, converted)

}
func decodeAndConvert(addrStr string) (string, []byte, error) {
	hrp, data, err := bech32.Decode(addrStr)
	if err != nil {
		return "", nil, fmt.Errorf("decode bech32 failed:%s", err)
	}
	converted, err := bech32.ConvertBits(data, 5, 8, false)
	if err != nil {
		return "", nil, fmt.Errorf("convert bits failed:%s", err)
	}
	return hrp, converted, nil
}

func IsFedAddress(s string) bool {
	hrp, bytes, err := decodeAndConvert(s)
	if err != nil || (hrp != ChordAddressHRP) || len(bytes) != AddressLength {
		return false
	}
	return true
}
func HexToAddress(s string) (addr Address, err error) {
	hrp, bytes, err := decodeAndConvert(s)
	if err != nil || (hrp != ChordAddressHRP) || len(bytes) != AddressLength {
		return addr, err
	}
	return BytesToAddress(bytes), nil
}

func (a Address) Hex() string {
	str, err := convertAndEncode(ChordAddressHRP, a[:])
	if err != nil {
		return ""
	}
	return str
}

func (a Address) String() string {
	return a.Hex()
}

func (a Address) MarshalText() ([]byte, error) {
	//result := make([]byte, len(a[:])*2+3)
	//copy(result, []byte("fed"))
	//hex.Encode(result[3:], a[:])
	return []byte(a.Hex()), nil
}

func (a *Address) UnmarshalBinary(data []byte) error {
	panic("22")
}

func (a *Address) UnmarshalText(input []byte) error {
	panic("2")
}

func (a *Address) UnmarshalJSON(input []byte) error {
	fmt.Println(input)
	fmt.Println(string(input))
	addr, err := HexToAddress(string(input))
	if err != nil {
		return err
	}
	a.SetBytes(addr[:])
	return nil
}

// Scan implements Scanner for database/sql.
func (a *Address) Scan(src interface{}) error {
	panic("4")
}

// Value implements valuer for database/sql.
func (a Address) Value() (driver.Value, error) {
	panic("5")
}

func (a Address) Format(s fmt.State, c rune) {
	panic("6")
}

func BytesToAddress(b []byte) Address {
	var a Address
	a.SetBytes(b)
	return a
}

func InterfaceToAddress(b string) (Address, error) {
	return HexToAddress(b)
}

func PubKeyToAddr(p *bls.PublicKey) Address {
	pubBytes := p.Serialize()
	return BytesToAddress(Keccak256(pubBytes[1:])[12:])
}
