package wallet

import (
	"encoding/hex"
	"encoding/json"
	"github.com/herumi/bls-eth-go-binary/bls"
	"github.com/otcChain/chord-go/common"
	"github.com/pborman/uuid"
)

const (
	version   = 1
	PriKeyLen = 32
)

type Key struct {
	ID         uuid.UUID
	Light      bool
	Address    common.Address
	PrivateKey *bls.SecretKey
}

type encryptedKeyJSON struct {
	Address string     `json:"address"`
	Crypto  CryptoJSON `json:"crypto"`
	ID      string     `json:"id"`
	Version int        `json:"version"`
}

type CryptoJSON struct {
	Cipher       string                 `json:"cipher"`
	CipherText   string                 `json:"ciphertext"`
	CipherParams cipherParamsJSON       `json:"cipherParams"`
	KDF          string                 `json:"kdf"`
	KDFParams    map[string]interface{} `json:"kdfParams"`
	MAC          string                 `json:"mac"`
}
type cipherParamsJSON struct {
	IV string `json:"iv"`
}

type plainKeyJSON struct {
	Address    string `json:"address"`
	PrivateKey string `json:"privateKey"`
	ID         string `json:"id"`
	Version    int    `json:"version"`
}

func (k *Key) MarshalJSON() (j []byte, err error) {
	jStruct := plainKeyJSON{
		hex.EncodeToString(k.Address[:]),
		hex.EncodeToString(k.PrivateKey.Serialize()),
		k.ID.String(),
		version,
	}
	j, err = json.Marshal(jStruct)
	return j, err
}

// UnmarshalJSON will unmarshal a byte array to the Key object
func (k *Key) UnmarshalJSON(j []byte) (err error) {
	keyJSON := new(plainKeyJSON)
	err = json.Unmarshal(j, &keyJSON)
	if err != nil {
		return err
	}

	u := new(uuid.UUID)
	*u = uuid.Parse(keyJSON.ID)
	k.ID = *u

	var sec bls.SecretKey
	secByte, err := hex.DecodeString(keyJSON.PrivateKey)
	if err != nil {
		return err
	}
	if err = sec.Deserialize(secByte); err != nil {
		return err
	}
	k.PrivateKey = &sec
	pub := sec.GetPublicKey()
	k.Address = common.PubKeyToAddr(pub)
	return nil
}

func NewKey() *Key {
	return NewLightKey(false)
}

func NewLightKey(light bool) *Key {
	sec := GenerateKey()
	id := uuid.NewRandom()
	key := &Key{
		Light:      light,
		ID:         id,
		Address:    common.PubKeyToAddr(sec.GetPublicKey()),
		PrivateKey: sec,
	}
	return key
}

func (k *Key) Encrypt(auth string) ([]byte, error) {
	if k.Light {
		return EncryptKey(k, auth, LightScryptN, LightScryptP)
	} else {
		return EncryptKey(k, auth, StandardScryptN, StandardScryptP)
	}
}

func (k *Key) isOpen() bool {
	return k.PrivateKey == nil
}

func (k *Key) close() {
	k.PrivateKey = nil
}

func GenerateKey() *bls.SecretKey {
	var sec bls.SecretKey
	sec.SetByCSPRNG()
	return &sec
}
