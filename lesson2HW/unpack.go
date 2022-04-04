package lesson2

import (
	"strconv"
	"strings"
	"unicode"
)

func Unpack(str string) string {
	_, err := strconv.Atoi(str)
	if err == nil {
		return ""
	} else {
		var output strings.Builder
		var prev rune
		for idx, char := range str {

			if unicode.IsDigit(char) && idx != 0 {
				n := int(char - '0')
				output.WriteString(strings.Repeat(strconv.QuoteRune(prev), n-1))

			} else {
				output.WriteRune(char)
			}
			prev = char
		}
		return strings.Join(strings.Split(output.String(), "'"), "")

	}
}
