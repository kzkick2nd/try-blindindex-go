package cmd

import (
	"github.com/spf13/cobra"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		if err := initAction(); err != nil {
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

type Entity struct {
	ID         int    `db:"id"`
	Entity     []byte `db:"entity"`
	EntityBidx []byte `db:"entity_bidx"`
}

func initAction() (err error) {
	db, _ := sqlx.Connect("sqlite3", "__sqlite.db")
	db.MustExec(schema)

	return nil
}
