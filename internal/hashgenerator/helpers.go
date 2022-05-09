package hashgenerator

import (
	"crypto/sha1"
	"fmt"
)

const salt = "dsadasjfds8f7y8hsouihfasd"

func GeneratePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
