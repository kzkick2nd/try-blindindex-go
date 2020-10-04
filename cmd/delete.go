package cmd

import (
	"fmt"
	"strconv"
	"github.com/spf13/cobra"

	"../lib"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		if err := deleteAction(args); err != nil {
			Exit(err, 1)
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}

func deleteAction(args []string) (err error) {
	fmt.Println("This is delete command")
	id, _ := strconv.Atoi(args[0])
	return blindindex.DeleteByID(id)
}
