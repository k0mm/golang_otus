package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("Invalid string!!!")

func nextRune(index int, slice []rune) rune {
	value := '1'
	if index < len(slice)-1 {
		value = slice[index+1]
	}
	return value
}

func Unpack(str string) (string, error) {
	if str == "" {
		return "", nil
	}

	runes := []rune(str)
	count := len(runes)
	var strBuilder strings.Builder

	for i := 0; i < count; i++ {
		symbol := runes[i]
		next := nextRune(i, runes)

		if unicode.IsDigit(symbol) {
			return "", ErrInvalidString
		}

		if symbol == '\\' {
			if !(next == '\\' || unicode.IsDigit(next)) || (next == '1' && i == count-1) {
				return "", ErrInvalidString
			}
			i++
			symbol = next
			next = nextRune(i, runes)
		}

		if unicode.IsDigit(next) {
			i++
		} else {
			next = '1'
		}

		repeat, _ := strconv.Atoi(string(next))
		strBuilder.WriteString(strings.Repeat(string(symbol), repeat))
	}

	return strBuilder.String(), nil
}
