package utils

import "strings"

// List of offensive words to filter (expand this list as needed)
var offensiveWords = []string{"fuck", "shit", "damn"}

// ProfanityFilter replaces offensive words with asterisks
func ProfanityFilter(message string) string {
	sanitizedMessage := message
	for _, word := range offensiveWords {
		if strings.Contains(strings.ToLower(sanitizedMessage), strings.ToLower(word)) {
			sanitizedMessage = strings.ReplaceAll(sanitizedMessage, word, "****")
		}
	}
	return sanitizedMessage
}
