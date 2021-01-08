package main

import (
	"fmt"

	"github.com/kawabatas/go-util/securerandom"
)

func main() {
	// securerandom.Hex
	fmt.Println(securerandom.Hex(16))
}
