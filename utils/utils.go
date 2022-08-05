package utils

import (
	"crypto/rand"
	"encoding/base64"
)

// ------------------------------------------------------------

func GenerateRandomStrings(n int) (string, error) {
	b, err := GenerateRandomBytes(n)
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(b), nil

}

// ------------------------------------------------------------

func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// ------------------------------------------------------------
