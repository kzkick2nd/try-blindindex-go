package cmd

import (
	"fmt"
	"github.com/spf13/cobra"

	"../lib"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var findCmd = &cobra.Command{
	Use:   "find",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		if err := findAction(args); err != nil {
			Exit(err, 1)
		}
	},
}

func init() {
	rootCmd.AddCommand(findCmd)
}

func findAction(args []string) (err error) {
	fmt.Println("This is find command")

	findByEntity := []Entity{}
	key, _ := encryption.CalcBlindIndex([]byte(salt), []byte(args[0]), truncate)

	db, _ := sqlx.Connect("sqlite3", "__sqlite.db")
	db.Select(&findByEntity, "SELECT * FROM entities WHERE entity_bidx=$1", key)
	for _, v := range findByEntity {
		plainText, _ := encryption.DecryptByGCM(encryptionKey, v.Entity)
		fmt.Println(v.ID, plainText, v.EntityBidx)
	}

	return nil
}
