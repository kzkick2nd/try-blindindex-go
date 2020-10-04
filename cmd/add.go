package cmd

import (
	"fmt"
	"github.com/spf13/cobra"

	"../lib"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		if err := addAction(); err != nil {
			Exit(err, 1)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

func addAction() (err error) {
	fmt.Println("This is add command")

	cipherName, _ := encryption.EncryptByGCM(encryptionKey, "有賀和輝")
	hashedName, _ := encryption.CalcBlindIndex([]byte(salt), []byte("有賀和輝"), truncate)

	db, _ := sqlx.Connect("sqlite3", "__sqlite.db")
	tx := db.MustBegin()
	tx.NamedExec(
		"INSERT INTO entities (entity, entity_bidx) VALUES (:entity, :entity_bidx)",
		&Entity{
			Entity:     cipherName,
			EntityBidx: hashedName,
		})
	tx.Commit()

	return nil
}
