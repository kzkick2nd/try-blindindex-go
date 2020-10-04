package cmd

import (
	"fmt"
	"github.com/spf13/cobra"

	"../lib"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		if err := addAction(args); err != nil {
			Exit(err, 1)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

func addAction(args []string) (err error) {
	if err := blindindex.SaveWithBlindIndex(args[0]); err != nil {
		return err
	}
	fmt.Println("Create record successfully")
	return nil
}
