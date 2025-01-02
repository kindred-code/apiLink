package connections

import (
	"crypto/sha256"
	"encoding/hex"
)

func shaHashing(input string) string {
	plainText := []byte(input)
	sha256Hash := sha256.Sum256(plainText)
	return hex.EncodeToString(sha256Hash[:])
}
