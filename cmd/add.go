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
	fmt.Println("this is add command")
},
}

func init() {
  rootCmd.AddCommand(addCmd)
}
