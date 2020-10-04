package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
)

var deleteCmd = &cobra.Command{
  Use:   "delete",
  Short: "",
  Long: "",
  Run: func(cmd *cobra.Command, args []string) {
    if err := deleteAction(); err != nil {
        Exit(err, 1)
    }
  },
}

func init() {
  rootCmd.AddCommand(deleteCmd)
}

func deleteAction() (err error) {
  fmt.Println("This is delete command")
  return nil
}