package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	h := sha256.New()
	s := "salt"
	p := "password"
	i := s + p

	h.Write([]byte(i))
	fmt.Println(hex.EncodeToString(h.Sum(nil)))

	hash, _ := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
	fmt.Println(string(hash))
}