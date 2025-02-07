package utils

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

func MustNewId() string {
	n := 8

	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		panic(err)
	}

	id := make([]byte, n*2)
	hex.Encode(id, bytes)

	return fmt.Sprintf("0x%s", id)
}
