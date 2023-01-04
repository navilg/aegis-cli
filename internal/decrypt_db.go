package internal

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
	"fmt"

	"golang.org/x/crypto/scrypt"
)

func DecryptDB(vaultdata Vault) ([]byte, error) {

	password := PasswordTUI()

	header := vaultdata.Header

	// Try the given password on every slot until one succeeds
	var masterkey []byte = nil
	for _, slot := range header.Slots {
		if slot.Type != 1 {
			continue
		}

		// Derive a key from password
		salt, _ := hex.DecodeString(slot.Salt)
		key, err := scrypt.Key(password, salt, slot.N, slot.R, slot.P, 32)
		if err != nil {
			// fmt.Println(1, err.Error())
			return nil, err
		}

		// Try to use derived key to decrypt master key
		encrypted_master_key, _ := hex.DecodeString(slot.Key)
		nonce, _ := hex.DecodeString(slot.KeyParams.Nonce)
		tag, _ := hex.DecodeString(slot.KeyParams.Tag)

		// ciphertext := []byte(string(encrypted_master_key) + string(tag))
		// fmt.Println(ciphertext)
		ciphertext := append(encrypted_master_key, tag...)

		block, err := aes.NewCipher(key)
		if err != nil {
			// fmt.Println(2, err.Error())
			return nil, err
		}
		aesgcm, err := cipher.NewGCM(block)
		if err != nil {
			// fmt.Println(3, err.Error())
			return nil, err
		}
		masterkey, err = aesgcm.Open(nil, nonce, ciphertext, nil)
		if err != nil {
			// fmt.Println(4, err.Error())
			continue
		}
		break
	}

	if masterkey == nil {
		return nil, fmt.Errorf("error: Unable to authenticate using provided password.")
	}
	encrypted_db, err := base64.StdEncoding.DecodeString(vaultdata.DB)
	if err != nil {
		return nil, err
	}

	nonce, _ := hex.DecodeString(vaultdata.Header.Params.Nonce)
	tag, _ := hex.DecodeString(vaultdata.Header.Params.Tag)
	ciphertext := []byte(string(encrypted_db) + string(tag))

	block, err := aes.NewCipher(masterkey)
	if err != nil {
		return nil, err
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	data, err := aesgcm.Open(nil, nonce, ciphertext, nil)

	return data, nil
}
