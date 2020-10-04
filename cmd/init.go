package cmd

import (
	"github.com/spf13/cobra"
	"fmt"

	"encoding/hex"

	"crypto/sha256"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"golang.org/x/crypto/pbkdf2"

	_ "github.com/mattn/go-sqlite3"
	"github.com/jmoiron/sqlx"
)

var initCmd = &cobra.Command{
  Use:   "init",
  Short: "",
  Long: "",
  Run: func(cmd *cobra.Command, args []string) {
    if err := addAction(); err != nil {
        Exit(err, 1)
    }
  },
}

func init() {
  rootCmd.AddCommand(initCmd)
}

var schema = `
DROP TABLE IF EXISTS entities;
CREATE TABLE entities (
	id          INTEGER PRIMARY KEY,
  entity      VARCHAR(80)  DEFAULT '',
  entity_bidx VARCHAR(80)  DEFAULT ''
);
`

var salt = "/k@R5S#(7iN)vzDkaUH_>v-r@C. da|Yxh`X>}w$Q6-@&3z!^&|umH^8doJv&R;}"
var encryptionKey, _ = hex.DecodeString("6368616e676520746869732070617373776f726420746f206120736563726574")
var truncate = 16

type Entity struct {
	ID         int    `db:"id"`
	Entity     []byte `db:"entity"`
	EntityBidx []byte `db:"entity_bidx"`
}

func addAction() (err error) {
  cipherName, _ := encryptByGCM(encryptionKey, "有賀和輝")
	hashedName := pbkdf2.Key([]byte("有賀和輝"), []byte(salt), 1024, truncate, sha256.New)

  db, _ := sqlx.Connect("sqlite3", "__sqlite.db")
	db.MustExec(schema)

	tx := db.MustBegin()
	tx.NamedExec(
		"INSERT INTO entities (entity, entity_bidx) VALUES (:entity, :entity_bidx)",
		&Entity{
			Entity: cipherName,
			EntityBidx: hashedName,
		})
	tx.Commit()

 	selectAll := []Entity{}
	db.Select(&selectAll, "SELECT * FROM entities ORDER BY id ASC")
	for _, v := range selectAll {
		fmt.Printf("%+v\n",v)
	}

	query := pbkdf2.Key([]byte("有賀和輝"), []byte(salt), 1024, 16, sha256.New)
	findByEntity := []Entity{}
	db.Select(&findByEntity, "SELECT * FROM entities WHERE entity_bidx=$1", query)
	d, _ := decryptByGCM(encryptionKey, findByEntity[0].Entity)
  fmt.Println(d)

  return nil
}

func encryptByGCM(encryptionKey []byte, plainText string) ([]byte, error) {
    block, err := aes.NewCipher(encryptionKey); if err != nil {
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