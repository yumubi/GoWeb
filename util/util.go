package util

import (
	"math/rand"
)

func RandomString(i int) string {
	letters := []byte("abcdefghijklmnopqrstuvwxyz")
	result := make([]byte, i)
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}
