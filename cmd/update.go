package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strconv"

	"../lib"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		if err := updateAction(args); err != nil {
			Exit(err, 1)
		}
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}

func updateAction(args []string) (err error) {
	// TODO flagの追加
	id, _ := strconv.Atoi(args[0])
	if err := blindindex.UpdateByID(id, args[1]); err != nil {
		return err
	}
	fmt.Println("Update record successfully")
	return nil
}
