package utils

import (
	"crypto/sha1"
	"fmt"
)

func GenerateHash(data string) string {
	hash := sha1.New()
	hash.Write([]byte(data))
	return fmt.Sprintf("%x", hash.Sum(nil))
}
