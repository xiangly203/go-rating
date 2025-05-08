package utils

import (
	"crypto/rand"
	"math/big"
	"regexp"
)

func IsMobile(number string) bool {
	return regexp.MustCompile(`^1\d{10}$`).MatchString(number)
}

func GenCode(phone string) (string, error) {
	numbers := "0123456789"
	length := 6
	result := make([]byte, length)
	maxNum := big.NewInt(int64(len(numbers)))

	for i := range result {
		var index *big.Int
		index, err := rand.Int(rand.Reader, maxNum)
		if err != nil {
			return "", err
		}
		result[i] = numbers[index.Int64()]
	}
	return string(result), nil
}
