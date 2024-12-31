package main

import (
	"crypto/sha256"
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

func Rand(length int) (string, error) {
	if length <= 0 {
		return "", fmt.Errorf("length must be greater than 0")
	}

	hash := sha256.New()

	buf := make([]byte, length)
	_, err := rand.Read(buf)
	
	if err != nil {
		return "", err
	}

	hash.Write(buf)

	hashString := hex.EncodeToString(hash.Sum(nil))

	if length > len(hashString) {
		return "", fmt.Errorf("length is too large for the hash output")
	}

	return hashString[:length], nil
}

func main() {
	randomString, err := Rand(16)
	
	if err != nil {
		fmt.Println("Error generating random string:", err)
		return
	}

	fmt.Println(randomString)
}
