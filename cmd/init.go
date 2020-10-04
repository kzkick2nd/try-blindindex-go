package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
)

var initCmd = &cobra.Command{
  Use:   "init",
  Short: "",
  Long: "",
  Run: func(cmd *cobra.Command, args []string) {
	fmt.Println("this is init command")
},
}

func init() {
  rootCmd.initCommand(addCmd)
}
