package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
)

var updateCmd = &cobra.Command{
  Use:   "update",
  Short: "",
  Long: "",
  Run: func(cmd *cobra.Command, args []string) {
	fmt.Println("this is update command")
},
}

func init() {
  rootCmd.AddCommand(updateCmd)
}
