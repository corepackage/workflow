package cryptography

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"

	"github.com/corepackage/workflow/internal/constants"
	fileops "github.com/corepackage/workflow/pkg/fileops"
)

// Decrypt : To decrypt an encrypted file from a specified
func Decrypt(file string) ([]byte, error) {
	// Invoke file read function
	ciphertext, readErr := fileops.ReadFromFile(file)
	if readErr != nil {
		return nil, fmt.Errorf("Error while reading: %v", readErr)
	}

	c, err := aes.NewCipher([]byte(constants.ENC_DEC_KEY))
	if err != nil {
		return nil, fmt.Errorf("Error while generating cipher instance: %v", err)
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, fmt.Errorf("Error in GCM: %v", err)
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, fmt.Errorf("Error in Nonce: %v", err)
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, fmt.Errorf("Error in GCM Open: %v", err)
	}
	return plaintext, nil
}

// Encrypt : To encrypt file and store in a path
func Encrypt(file, writePath string) error {

	// Invoke file read function
	readString, readErr := fileops.ReadFromFile(file)
	if readErr != nil {
		return fmt.Errorf("Error while reading: %v", readErr)
	}
	// generate a new aes cipher using our 32 byte long key
	c, err := aes.NewCipher([]byte(constants.ENC_DEC_KEY))
	if err != nil {
		return fmt.Errorf("Error while generating cipher instance: %v", err)

	}

	// gcm or Galois/Counter Mode, is a mode of operation for symmetric key cryptographic block ciphers
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return fmt.Errorf("Error in GCM: %v", err)

	}

	// creates a new byte array the size of the nonce which must be passed to Seal
	nonce := make([]byte, gcm.NonceSize())
	// populates our nonce with a cryptographically secure
	// random sequence
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return fmt.Errorf("Error while reading nonce: %v", err)

	}

	// here we encrypt our text using the Seal function
	// Seal encrypts and authenticates plaintext, authenticates the
	// additional data and appends the result to dst, returning the updated
	// slice. The nonce must be NonceSize() bytes long and unique for all
	// time, for a given key.

	writeErr := fileops.WriteToFile(string(gcm.Seal(nonce, nonce, readString, nil)), writePath)
	if writeErr != nil {
		return fmt.Errorf("Error while writing: %v", writeErr)

	}

	return nil
}
