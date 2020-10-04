package main

import (
	"./cmd"

	"crypto/sha256"
	"encoding/hex"
	"fmt"
	// "log"
	"golang.org/x/crypto/pbkdf2"
	// "golang.org/x/crypto/bcrypt"

	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"

	// "database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/jmoiron/sqlx"
)

var schema = `
DROP TABLE IF EXISTS user;
CREATE TABLE user (
	id     INTEGER PRIMARY KEY,
    name   VARCHAR(80)  DEFAULT '',
    namebi VARCHAR(80)  DEFAULT ''
);
`

type User struct {
	ID      int    `db:"id"`
	Name    []byte `db:"name"`
	NameBI  []byte `db:"namebi"`
}

func main() {
	cmd.Execute()
	salt := "salt"
	key, _ := hex.DecodeString("6368616e676520746869732070617373776f726420746f206120736563726574")

	// pbkdf2 stretch 2**10
	dk := pbkdf2.Key([]byte("arugakazuki@lifull.com"), []byte(salt), 1024, 8, sha256.New)
	fmt.Println(hex.EncodeToString(dk))

	// AES GCM
    // cipherText, _ := encryptByGCM(key, "12345")
	// fmt.Println(hex.EncodeToString(cipherText))
	// decryptedText, _ := decryptByGCM(k, cipherText)
    // fmt.Printf("Decrypted Text: %v\n ", decryptedText)

	// sqlite
	db, _ := sqlx.Connect("sqlite3", "__sqlite.db")
	db.MustExec(schema)

	cipherName, _ := encryptByGCM(key, "有賀和輝")
	hashedName := pbkdf2.Key([]byte("有賀和輝"), []byte(salt), 1024, 16, sha256.New)

	tx := db.MustBegin()
	tx.NamedExec(
		"INSERT INTO user (Name, NameBI) VALUES (:name, :namebi)",
		&User{
			Name: cipherName,
			NameBI: hashedName,
		})
	tx.Commit()

	selectAll := []User{}
	db.Select(&selectAll, "SELECT * FROM user ORDER BY id ASC")
	for _, v := range selectAll {
		fmt.Printf("%+v\n",v)
	}

	query := pbkdf2.Key([]byte("有賀和輝"), []byte(salt), 1024, 16, sha256.New)
	findByName := []User{}
	db.Select(&findByName, "SELECT * FROM user WHERE namebi=$1", query)
	d, _ := decryptByGCM(key, findByName[0].Name)
	fmt.Println(d)
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