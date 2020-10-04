package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
)

var findCmd = &cobra.Command{
  Use:   "find",
  Short: "",
  Long: "",
  Run: func(cmd *cobra.Command, args []string) {
    if err := findAction(); err != nil {
        Exit(err, 1)
    }
  },
}

func init() {
  rootCmd.AddCommand(findCmd)
}

func findAction() (err error) {
  fmt.Println("This is find command")
  return nil
}