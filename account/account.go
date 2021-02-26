package account

import (
	"encoding/json"
	"io/ioutil"
)

const ConfFileName = "account.json"

type Account struct {
}

func NewAccount(dir, password string) (err error) {
	acc := &Account{}
	bts, e := json.MarshalIndent(acc, "", "\t")
	if e != nil {
		return e
	}
	if err = ioutil.WriteFile(dir+"/"+ConfFileName, bts, 0644); err != nil {
		return
	}
	return nil
}
