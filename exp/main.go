package main

import (
	"fmt"
	"webdev/hash"
)

func main() {
	//toHash := []byte("This is the text")
	//h := hmac.New(sha256.New, []byte("My secret Key"))
	//h.Write(toHash)
	//b := h.Sum(nil)
	//fmt.Println(b)

	hmac := hash.NewHMAC("my-secret-key")
	fmt.Println(hmac.Hash("This is the text"))
}
