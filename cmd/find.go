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
	fmt.Println("this is find command")
},
}

func init() {
  rootCmd.AddCommand(findCmd)
}
