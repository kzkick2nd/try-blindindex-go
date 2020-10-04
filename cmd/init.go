package cmd

import (
	"fmt"
	"github.com/spf13/cobra"

	"../lib"
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

func initAction() (err error) {
	if err := blindindex.InitTable(); err != nil {
		return err
	}
	fmt.Println("Init sqlite3 DB successfully")
	return nil
}
