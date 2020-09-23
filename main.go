package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/pbkdf2"

	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"

	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/jmoiron/sqlx"
)

var schema = `
DROP TABLE IF EXISTS user;
CREATE TABLE user (
	user_id    INTEGER PRIMARY KEY,
    first_name VARCHAR(80)  DEFAULT '',
    last_name  VARCHAR(80)  DEFAULT '',
	email      VARCHAR(250) DEFAULT '',
	password   VARCHAR(250) DEFAULT NULL
);
`

type User struct {
	UserID    int    `db:"user_id"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Email     string
	Password  sql.NullString
}

func main() {
	s := "salt"
	p := "password"

	// pbkdf2 stretch 2**10
	dk := pbkdf2.Key([]byte(p), []byte(s), 1024, 32, sha256.New)
	fmt.Println(hex.EncodeToString(dk))

	// bcrypt
	hash, _ := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
	fmt.Println(string(hash))

	// AES GCM
	key, _ := hex.DecodeString("6368616e676520746869732070617373776f726420746f206120736563726574")
    cipherText, _ := encryptByGCM(key, "12345")
	fmt.Println(hex.EncodeToString(cipherText))
    decryptedText, _ := decryptByGCM(key, cipherText)
    fmt.Printf("Decrypted Text: %v\n ", decryptedText)

	// sqlite
	db, _ := sqlx.Connect("sqlite3", "__deleteme.db")
	db.MustExec(schema)
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