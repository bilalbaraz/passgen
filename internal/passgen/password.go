package passgen

import (
	"crypto/rand"
	"math/big"
	"strings"
)

func GeneratePassword(length int, charset []rune) (string, error) {
	var b strings.Builder
	b.Grow(length)
	max := big.NewInt(int64(len(charset)))

	for i := 0; i < length; i++ {
		idx, err := rand.Int(rand.Reader, max)
		if err != nil {
			return "", err
		}
		b.WriteRune(charset[idx.Int64()])
	}
	return b.String(), nil
}
