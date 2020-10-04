package blindindex

import (
	"fmt"
	"encoding/hex"

	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"

	"golang.org/x/crypto/pbkdf2"
	"crypto/sha256"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var salt = "/k@R5S#(7iN)vzDkaUH_>v-r@C. da|Yxh`X>}w$Q6-@&3z!^&|umH^8doJv&R;}"
var encryptionKey, _ = hex.DecodeString("6368616e676520746869732070617373776f726420746f206120736563726574")
var truncate = 16

type Entity struct {
	ID         int    `db:"id"`
	Entity     []byte `db:"entity"`
	EntityBidx []byte `db:"entity_bidx"`
}

func InitTable() (err error){
	schema := `
		DROP TABLE IF EXISTS entities;
		CREATE TABLE entities (
			id          INTEGER PRIMARY KEY,
			entity      VARCHAR(80)  DEFAULT '',
			entity_bidx VARCHAR(80)  DEFAULT ''
		);
	`
	db, _ := sqlx.Connect("sqlite3", "__sqlite.db")
	db.MustExec(schema)

	return nil
}

func SaveWithBlindIndex(plainText string) (err error){
	cipherText, _ := encryptByGCM(encryptionKey, plainText)
	blindIndex, _ := calcBlindIndex([]byte(salt), []byte(plainText), truncate)

	db, _ := sqlx.Connect("sqlite3", "__sqlite.db")
	tx := db.MustBegin()
	tx.NamedExec(
		"INSERT INTO entities (entity, entity_bidx) VALUES (:entity, :entity_bidx)",
		&Entity{
			Entity:     cipherText,
			EntityBidx: blindIndex,
		})
	tx.Commit()
	// TODO 追加レコード内容の表示

	return nil
}

func FindHumanByPlainText(plainText string) (err error){
	findByEntity := []Entity{}
	key, _ := calcBlindIndex([]byte(salt), []byte(plainText), truncate)

	db, _ := sqlx.Connect("sqlite3", "__sqlite.db")
	db.Select(&findByEntity, "SELECT * FROM entities WHERE entity_bidx=$1", key)
	for _, v := range findByEntity {
		plainText, _ := decryptByGCM(encryptionKey, v.Entity)
		fmt.Println(v.ID, plainText, v.EntityBidx)
	}
	// TODO ここにフィルター追加
	// TODO テーブル構造で出力
	return nil
}

// func UpdateByID(){}

// func DeleteByID(){}

// func ShowRawTable(){}

func calcBlindIndex(salt []byte, plainText []byte, keyLen int) ([]byte, error) {
	return pbkdf2.Key(plainText, salt, 1024, keyLen, sha256.New), nil
}

func encryptByGCM(encryptionKey []byte, plainText string) ([]byte, error) {
	block, _ := aes.NewCipher(encryptionKey)
	gcm, _ := cipher.NewGCM(block)

	nonce := make([]byte, gcm.NonceSize())
	_, _ = rand.Read(nonce)

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
