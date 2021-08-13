package app

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"regexp"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	tokenLength := 16
	tokenSymbols := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	n := 10000
	var token string
	tokenRegex := fmt.Sprintf("^[a-zA-Z0-9]{%d}$", tokenLength)

	for i := 0; i < n; i++ {
		token = GenerateToken(tokenLength, tokenSymbols)
		isValid, err := regexp.MatchString(tokenRegex, token)
		if err != nil {
			log.Println(err)
		}
		assert.True(t, isValid)
	}
}

func TestTokenGenerationUniqueness(t *testing.T) {
	tokenLength := 16
	tokenSymbols := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	n := 10000
	var token string
	var isUnique bool
	tokens := make([]string, 0, n)

	for i := 0; i < n; i++ {
		token = GenerateToken(tokenLength, tokenSymbols)

		isUnique = !contains(&tokens, token)

		assert.True(t, isUnique)
		tokens = append(tokens, token)
	}
}

//func contains(s *[]string, e string) bool {
//	for i := range *s {
//		if (*s)[i] == e {
//			return true
//		}
//	}
//	return false
//}
