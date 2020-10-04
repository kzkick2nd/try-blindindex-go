package cmd

import (
	"fmt"
	"encoding/hex"

	"github.com/spf13/cobra"
	"os"
)

var salt = "/k@R5S#(7iN)vzDkaUH_>v-r@C. da|Yxh`X>}w$Q6-@&3z!^&|umH^8doJv&R;}"
var encryptionKey, _ = hex.DecodeString("6368616e676520746869732070617373776f726420746f206120736563726574")
var truncate = 16

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
