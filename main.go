package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	h := sha256.New()
	s := "salt"
	k := "password"
	i := s + k

	h.Write([]byte(i))
	fmt.Printf("%x", h.Sum(nil)[:2])
}