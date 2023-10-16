package random

import "math/rand"

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789"
const length = 14

// ".zip"
const format = ".zip"
const formatLength = len(format)

func RandomZipFileName() string {
	b := make([]byte, length+formatLength)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	b = append(b, []byte(format)...)

	return string(b)
}
