package cmd

import (
	"github.com/spf13/cobra"

	"../lib"
)

var selectCmd = &cobra.Command{
	Use:   "select-all",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		if err := selectAction(); err != nil {
			Exit(err, 1)
		}
	},
}

func init() {
	rootCmd.AddCommand(selectCmd)
}

func selectAction() (err error) {
	return blindindex.ShowRowTable()
}
