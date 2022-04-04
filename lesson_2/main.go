package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func unpackString(str string) {
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
	fmt.Println(strings.Join(strings.Split(output.String(), "'"), ""))

}

func main() {
	unpackString("a2b3//5t6")
}
