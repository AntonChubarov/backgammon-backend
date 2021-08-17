package auth

import (
	"backgammon/config"
	"backgammon/utils"
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"regexp"
	"testing"
)

var conf = config.ServerConfig{
		Token: config.TokenConfig{
			TokenLength: 16,
			TokenSymbols: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"},
			}
var generator = NewTokenGeneratorFlex(&conf)

func TestGenerateToken(t *testing.T) {

	n := 10000
	var token string
	tokenRegex := fmt.Sprintf("^[a-zA-Z0-9]{%d}$", conf.Token.TokenLength)

	for i := 0; i < n; i++ {
		token= generator.GenerateToken()
		isValid, err := regexp.MatchString(tokenRegex, token)
		if err != nil {
			log.Println(err)
		}
		assert.True(t, isValid)
	}
}

func TestTokenGenerationUniqueness(t *testing.T) {

	n := 10000
	var token string
	var isUnique bool
	tokens := make([]string, 0, n)

	for i := 0; i < n; i++ {
		token = generator.GenerateToken()

		isUnique = !utils.Contains(&tokens, token)

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