package str

import (
	"strings"
	"unicode"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// EmptyString the empty string ""
const EmptyString = ""

func LowerFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return EmptyString
}

func UpperFirst(str string) string {
	return cases.Title(language.Und).String(str)
}

// ToLowerCamel converts a string to lowerCamelCase
func ToLowerCamel(s string) string {
	if s == "" {
		return s
	}
	if r := rune(s[0]); r >= 'A' && r <= 'Z' {
		s = strings.ToLower(string(r)) + s[1:]
	}
	return toCamelInitCase(s, false)
}
