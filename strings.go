package goutils

import (
	"strings"
	"unicode"
)

// ToCamelCase converts a snake_case string to camelCase.
func ToCamelCase(str string) string {
	words := strings.Split(str, "_")
	for i := range words {
		if i > 0 {
			// Manually capitalize the first letter and make the rest of the word lowercase
			words[i] = capitalizeFirstLetter(words[i])
		}
	}
	return strings.Join(words, "")
}

// capitalizeFirstLetter capitalizes the first letter of a word and lowercases the rest.
func capitalizeFirstLetter(word string) string {
	// If the word is empty, return it as-is
	if len(word) == 0 {
		return word
	}

	// Convert the first letter to uppercase and the rest to lowercase
	return string(unicode.ToUpper(rune(word[0]))) + word[1:]
}
