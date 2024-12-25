package hw02unpackstring

import (
	"errors"
	"strconv"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(input string) (string, error) {
	var result string
	runes := []rune(input)

	for i := 0; i < len(runes); i++ {
		current := runes[i]

		if !unicode.IsDigit(current) {
			result += string(current)
		} else {
			if i == 0 || unicode.IsDigit(rune(input[i-1])) {
				return "", ErrInvalidString
			}

			count, _ := strconv.Atoi(string(current))
			if count > 0 {
				for j := 0; j < count-1; j++ {
					result += string(input[i-1])
				}
			} else if count == 0 {
				result = result[:len(result)-1]
			}
		}
	}
	return result, nil
}
