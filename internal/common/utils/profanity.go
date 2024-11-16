package utils

import (
	"github.com/TwiN/go-away"
)

// ProfanityFilter replaces offensive words with asterisks using the go-away library
func ProfanityFilter(message string) string {
	sanitizedMessage := goaway.Censor(message)

	return sanitizedMessage
}
