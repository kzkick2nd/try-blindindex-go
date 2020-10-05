package blindindex

import (
	"encoding/hex"
	"fmt"
	"strconv"

	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"

	"crypto/sha256"
	"golang.org/x/crypto/pbkdf2"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"

	"os"
	_ "github.com/joho/godotenv/autoload"
)

var salt             = os.Getenv("SALT")
var encryptionKey, _ = hex.DecodeString(os.Getenv("ENCRYPTION_KEY"))
var truncate, _      = strconv.Atoi(os.Getenv("TRUNCATE"))

type Entity struct {
	ID       int    `db:"id"`
	Text     []byte `db:"text"`
	TextBidx []byte `db:"text_bidx"`
}

func InitTable() (err error) {
	schema := `
		DROP TABLE IF EXISTS entities;
		CREATE TABLE entities (
			id        INTEGER PRIMARY KEY,
			text      VARCHAR(80)  DEFAULT '',
			text_bidx VARCHAR(80)  DEFAULT ''
		);
	`
	db, _ := sqlx.Connect("sqlite3", "__sqlite.db")
	db.MustExec(schema)
	return nil
}

func SaveWithBlindIndex(plainText string) (err error) {
	cipherText, _ := encryptByGCM(encryptionKey, plainText)

	fieldName := "text"
	blindIndex, _ := calcBlindIndex([]byte(salt + fieldName), []byte(plainText), truncate)

	db, _ := sqlx.Connect("sqlite3", "__sqlite.db")
	tx := db.MustBegin()
	tx.NamedExec(
		"INSERT INTO entities (text, text_bidx) VALUES (:text, :text_bidx)",
		&Entity{
			Text:     cipherText,
			TextBidx: blindIndex,
		})
	tx.Commit()
	return nil
}

func FindHumanByPlainText(searchText string) (err error) {
	fieldName := "text"
	key, _ := calcBlindIndex([]byte(salt + fieldName), []byte(searchText), truncate)

	findByBidx := []Entity{}
	db, _ := sqlx.Connect("sqlite3", "__sqlite.db")
	db.Select(&findByBidx, "SELECT * FROM entities WHERE text_bidx=$1", key)

	fmt.Printf("Bidx matched: %v\n",len(findByBidx))
	findByPlainText := []Entity{}
	for _, v := range findByBidx {
		decryptedText, _ := decryptByGCM(encryptionKey, v.Text)
		if decryptedText == searchText {
			findByPlainText = append(findByPlainText, v)
		}
	}

	fmt.Printf("Results: %v\n",len(findByPlainText))
	for _, v := range findByPlainText {
		decryptedText, _ := decryptByGCM(encryptionKey, v.Text)
		fmt.Printf(
			"ID: %v\nDecrypted: %v\nSavedText: %v\nSavedBidx: %v\n\n",
			v.ID, decryptedText, v.Text, v.TextBidx)
	}
	return nil
}

func UpdateByID(ID int, plainText string) (err error) {
	cipherText, _ := encryptByGCM(encryptionKey, plainText)
	blindIndex, _ := calcBlindIndex([]byte(salt), []byte(plainText), truncate)

	db, _ := sqlx.Connect("sqlite3", "__sqlite.db")
	tx := db.MustBegin()
	tx.NamedExec(
		"UPDATE entities SET text=:text, text_bidx=:text_bidx WHERE id=:id",
		&Entity{
			ID:       ID,
			Text:     cipherText,
			TextBidx: blindIndex,
		})
	tx.Commit()
	return nil
}

func DeleteByID(ID int) (err error) {
	db, _ := sqlx.Connect("sqlite3", "__sqlite.db")
	db.Exec("DELETE FROM entities WHERE id=:id", ID)
	return nil
}

func ShowRowTable() (err error){
	selectAll := []Entity{}
	db, _ := sqlx.Connect("sqlite3", "__sqlite.db")
	db.Select(&selectAll, "SELECT * FROM entities ORDER BY id ASC")
	fmt.Printf("Records: %v\n",len(selectAll))
	for _, v := range selectAll {
		fmt.Printf("ID: %v\nEntity: %v\nBidx: %v\n\n",v.ID, v.Text, v.TextBidx)
	}
	return nil
}

func calcBlindIndex(salt []byte, plainText []byte, keyLen int) ([]byte, error) {
	return pbkdf2.Key(plainText, salt, 1024, keyLen, sha256.New), nil
}

func encryptByGCM(encryptionKey []byte, plainText string) ([]byte, error) {
	block, _ := aes.NewCipher(encryptionKey)
	gcm, _ := cipher.NewGCM(block)

	nonce := make([]byte, gcm.NonceSize())
	rand.Read(nonce)

	cipherText := gcm.Seal(nil, nonce, []byte(plainText), nil)
	cipherText = append(nonce, cipherText...)

	return cipherText, nil
}

func decryptByGCM(key []byte, cipherText []byte) (string, error) {
	block, _ := aes.NewCipher(key)
	gcm, _ := cipher.NewGCM(block)

	nonce := cipherText[:gcm.NonceSize()]
	plainByte, _ := gcm.Open(nil, nonce, cipherText[gcm.NonceSize():], nil)

	return string(plainByte), nil
}
