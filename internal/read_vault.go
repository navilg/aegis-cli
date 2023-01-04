package internal

import (
	"encoding/json"
	"io/ioutil"
)

func ReadVault(dbfilename string) (*Vault, error) {
	data, err := ioutil.ReadFile(dbfilename)
	if err != nil {
		return nil, err
	}

	var vault Vault
	json.Unmarshal(data, &vault)

	return &vault, nil
}
