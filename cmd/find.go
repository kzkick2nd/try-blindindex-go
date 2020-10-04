package cmd

import (
	"github.com/spf13/cobra"

	"../lib"
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
	return blindindex.FindHumanByPlainText(args[0])
}
