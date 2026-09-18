package namegenerator

import (
	"fmt"
	"math/rand/v2"
)

const charset = "abcdefghijklmnopqrstuvwxyz0123456789"

func GenerateName(prefix string) string {
	suffix := make([]byte, 5)
	for i := range suffix {
		suffix[i] = charset[rand.IntN(len(charset))]
	}
	return fmt.Sprintf("%s-%s", prefix, string(suffix))
}
