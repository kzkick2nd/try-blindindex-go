package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "blindindex-sample",
	Short: "hogehoge Short",
	Long:  "hogehoge Long",
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

func Exit(err error, codes ...int) {
	var code int
	if len(codes) > 0 {
		code = codes[0]
	} else {
		code = 2
	}
	if err != nil {
		fmt.Println(err)
	}
	os.Exit(code)
}
