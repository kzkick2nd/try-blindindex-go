package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strconv"

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
	id, _ := strconv.Atoi(args[0])
	if err := blindindex.DeleteByID(id); err != nil {
		return err
	}
	fmt.Println("Delete record successfully")
	return nil
}
