package hw02unpackstring

import (
	"errors"
	"strconv"
)

var ErrInvalidString = errors.New("invalid string")

func isSingleCharNumber(str string) bool {
	_, err := strconv.Atoi(str)
	return err == nil
}

func Unpack(input string) (string, error) {
	result := ""
	runes := []rune(input)
	var symbols []string

	for _, r := range runes {
		symbols = append(symbols, string(r))
	}
	length := len(symbols)

	for i := 0; i < length; i++ {
		current := symbols[i]

		if isSingleCharNumber(current) {
			if i == 0 || isSingleCharNumber(symbols[i-1]) {
				return "", ErrInvalidString
			}

			count, _ := strconv.Atoi(current)
			if count > 0 {
				for j := 0; j < count-1; j++ {
					result += symbols[i-1]
				}
			}
		} else {
			result += string(current)
		}
	}
	return result, nil
}
