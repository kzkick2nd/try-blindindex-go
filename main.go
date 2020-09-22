package main

import (
    "crypto/sha256"
    "encoding/hex"
	"fmt"
)

func main() {
	h := sha256.Sum256([]byte("Password"))
	fmt.Println(hex.EncodeToString(h[:]))
}