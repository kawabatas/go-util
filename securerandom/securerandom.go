package securerandom

import (
	"crypto/rand"
	"fmt"
)

// Hex is a port of ruby's SecureRandom.hex
func Hex(b int) string {
	k := make([]byte, b)
	if _, err := rand.Read(k); err != nil {
		panic(err)
	}
	return fmt.Sprintf("%x", k)
}
