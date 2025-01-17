package goutils

import (
	"crypto/rand"
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

// GenUpperCaseID generates a random ID consisting of uppercase letters.
// The length of the ID is determined by the input parameter 'length'.
// It uses the 'rand.Read' function to generate random values and maps them to
// the uppercase alphabet characters.
func GenUpperCaseID(length int) (string, error) {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var result []byte
	for i := 0; i < length; i++ {
		randomByte := make([]byte, 1)
		_, err := rand.Read(randomByte)
		if err != nil {
			return "", err
		}
		result = append(result, charset[int(randomByte[0])%len(charset)])
	}
	return string(result), nil
}
