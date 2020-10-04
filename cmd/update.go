package cmd

import (
	"fmt"
	"strconv"
	"github.com/spf13/cobra"

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
	fmt.Println("This is update command")
	// TODO flagの追加
	id, _ := strconv.Atoi(args[0])
	return blindindex.UpdateByID(id, args[1])
}
