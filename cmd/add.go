package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
)

var addCmd = &cobra.Command{
  Use:   "add",
  Short: "",
  Long: "",
  Run: func(cmd *cobra.Command, args []string) {
    if err := addAction(); err != nil {
        Exit(err, 1)
    }
  },
}

func init() {
  rootCmd.AddCommand(addCmd)
}

func addAction() (err error) {
  fmt.Println("This is add command")
  return nil
}