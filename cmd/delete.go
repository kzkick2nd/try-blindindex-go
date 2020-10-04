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
	fmt.Println("this is delete command")
},
}

func init() {
  rootCmd.AddCommand(deleteCmd)
}
