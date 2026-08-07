package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	runes := []rune(s)
	var sb strings.Builder

	for i := 0; i < len(runes); i++ {
		r := runes[i]

		if unicode.IsDigit(r) {
			return "", ErrInvalidString
		}

		counter := 1
		var err error

		if i+1 < len(runes) && unicode.IsDigit(runes[i+1]) {
			if i+2 < len(runes) && unicode.IsDigit(runes[i+2]) {
				return "", ErrInvalidString
			}
			counter, err = strconv.Atoi(string(runes[i+1]))
			if err != nil {
				return "", ErrInvalidString
			}

			i++
		}

		sb.WriteString(strings.Repeat(string(r), counter))
	}

	return sb.String(), nil
}
