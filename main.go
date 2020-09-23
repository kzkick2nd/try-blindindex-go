package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/bcrypt"

	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
)

func main() {
	s := "salt"
	p := "password"
	k := []byte(s + p)

	// stretch
	for i := 0; i < 1024; i++ {
		sum := sha256.Sum256(k)
		k = sum[:]
	}
	fmt.Println(hex.EncodeToString(k))

	hash, _ := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
	fmt.Println(string(hash))

	key, _ := hex.DecodeString("6368616e676520746869732070617373776f726420746f206120736563726574")
    cipherText, _ := encryptByGCM(key, "12345")
	fmt.Println(hex.EncodeToString(cipherText))
    decryptedText, _ := decryptByGCM(key, cipherText)
    fmt.Printf("Decrypted Text: %v\n ", decryptedText)

}

func encryptByGCM(key []byte, plainText string) ([]byte, error) {
    block, err := aes.NewCipher(key); if err != nil {
        return nil, err
    }

    gcm, err := cipher.NewGCM(block); if err != nil {
        return nil, err
    }

    nonce := make([]byte, gcm.NonceSize())// Unique nonce is required(NonceSize 12byte)
    _, err = rand.Read(nonce); if err != nil {
        return nil, err
    }

    cipherText := gcm.Seal(nil, nonce, []byte(plainText), nil)
    cipherText = append(nonce, cipherText...)

    return cipherText, nil
}

func decryptByGCM(key []byte, cipherText []byte) (string, error) {
    block, err := aes.NewCipher(key); if err != nil {
        return "", err
    }
    gcm, err := cipher.NewGCM(block); if err != nil {
        return "", err
    }

    nonce := cipherText[:gcm.NonceSize()]
    plainByte, err := gcm.Open(nil, nonce, cipherText[gcm.NonceSize():], nil); if err != nil {
        return "", err
    }

    return string(plainByte), nil
}