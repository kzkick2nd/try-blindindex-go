package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"os"
)

var rootCmd = &cobra.Command{
  Use:   "blindindex-sample",
  Short: "hogehoge Short",
  Long: "hogehoge Long",
  Run: func(cmd *cobra.Command, args []string) {
	fmt.Println("hello cobra")
},
}

func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Fprintln(os.Stderr, err)
    os.Exit(1)
  }
}