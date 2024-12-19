package hw02unpackstring

import (
	"errors"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(input string) (string, error) {
	result := ""
	length := len(input)

	for i := 0; i < length; i++ {
		current := input[i]

		if unicode.IsDigit(rune(current)) {
			if i == 0 || unicode.IsDigit(rune(input[i-1])) {
				return "", ErrInvalidString
			}

			count := int(current)
			if count > 0 {
				for j := 0; j < count-1; j++ {
					result += string(input[i-1])
				}
			}
		} else {
			result += string(current)
		}
	}
	return result, nil
}
