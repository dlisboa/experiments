package main

import (
	"crypto/rand"
	"fmt"
)

func main() {
	// rand.Text introduced in Go 1.24
	fmt.Printf("rand.Text(): %s\n", rand.Text())
}
