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
	fmt.Println("This is init command")
	return blindindex.InitTable()
}
