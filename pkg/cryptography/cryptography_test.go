package cryptography

import (
	"testing"
)

func TestEncrypt(t *testing.T) {
	cryptErr := Encrypt("./../../configs/configs.json", "./../../configs/enc-configs/testabc.json")
	if cryptErr != nil {
		t.Errorf("Error while encrypting file %v", cryptErr)
	}

}

func TestDecrypt(t *testing.T) {
	plainString, cryptErr := Decrypt("./../../configs/enc-configs/testabc.json")

	if cryptErr != nil {
		t.Errorf("Error while decrypting file %v", cryptErr)
	}
	t.Log("Decrypted String: ", plainString)
}
