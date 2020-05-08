package main

import (
	"fmt"
	"webdev/rand"
)

func main() {

	fmt.Println(rand.String(10))
	fmt.Println(rand.RememberToken())
}
