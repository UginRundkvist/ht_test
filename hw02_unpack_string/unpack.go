package hw02unpackstring

import (
	"errors"
	"strconv"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(input string) (string, error) {
	result := ""
	symbols := make([]string, 0)

	for _, r := range input {
		symbols = append(symbols, string(r))
	}

	for i := 0; i < len(symbols); i++ {
		current := symbols[i]

		if isStringNumber(current) {
			if i == 0 || isStringNumber(symbols[i-1]) {
				return "", ErrInvalidString
			}

			count, _ := strconv.Atoi(current)
			if count > 0 {
				for j := 0; j < count-1; j++ {
					result += symbols[i-1]
				}
			} else if count == 0 {
				result = result[:len(result)-1]
			}
		} else {

			result += current

		}
	}
	return result, nil
}

func isStringNumber(str string) bool {
	_, err := strconv.Atoi(str)
	return err == nil
}
