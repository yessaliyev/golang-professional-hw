package hw02unpackstring

import (
	"errors"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

const (
	numStart int32 = 48 // 0
	numEnd   int32 = 57 // 9
)

func Unpack(text string) (string, error) {
	if len(text) == 0 {
		return "", nil
	}

	start := int32(text[0])
	if start >= numStart && start <= numEnd {
		return "", ErrInvalidString
	}

	var b strings.Builder
	var tempChar string
	var checkNum bool

	for _, val := range text {
		if checkSymbol(val) {
			return "", ErrInvalidString
		}

		if val >= numStart && val <= numEnd {
			if checkNum {
				return "", ErrInvalidString
			}

			checkNum = true
			count := int(val - '0')

			if count == 0 {
				bText := b.String()
				newText := bText[:len(bText)-len(tempChar)]
				b.Reset()
				b.WriteString(newText)
				continue
			}

			tempChars := strings.Repeat(tempChar, count-1)
			b.WriteString(tempChars)
			tempChar = ""
		} else {
			tempChar = string(val)
			checkNum = false
			b.WriteString(tempChar)
		}
	}

	return b.String(), nil
}

func checkSymbol(val int32) bool {
	c1 := isEnglishOrCyrillicLetter(val)
	c2 := val >= numStart && val <= numEnd

	return !c1 && !c2
}

func isEnglishOrCyrillicLetter(r rune) bool {
	if unicode.Is(unicode.Latin, r) {
		return true
	}

	if unicode.Is(unicode.Cyrillic, r) {
		return true
	}

	// Тут можно чекать и другие типы

	return false
}
