package random

import (
	"crypto/rand"
	"encoding/hex"
)

func String(length int) (string, error) {
	bLength := length
	if bLength%2 != 0 {
		bLength += 1
	}
	bytes := make([]byte, bLength/2)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	randStr := hex.EncodeToString(bytes)
	if length != bLength {
		randStr = randStr[:len(randStr)-1]
	}
	return randStr, nil
}
