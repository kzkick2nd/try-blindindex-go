package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
)

var testCmd = &cobra.Command{
  Use:   "test",
  Short: "",
  Long: "",
  Run: func(cmd *cobra.Command, args []string) {
	fmt.Println("this is test command")
},
}

func init() {
  rootCmd.AddCommand(testCmd)
}
