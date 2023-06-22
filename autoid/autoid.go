package autoid

import (
	"crypto/rand"
	"fmt"
)

const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

func New() string {
	bytes := make([]byte, 20)

	if _, err := rand.Read(bytes); err != nil {
		panic(fmt.Sprintf("could not generate the id: %v", err))
	}

	for index, byt := range bytes {
		bytes[index] = alphabet[int(byt)%len(alphabet)]
	}

	return string(bytes)
}
