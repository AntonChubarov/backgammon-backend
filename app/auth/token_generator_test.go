package auth

import (
	"backgammon/config"
	"backgammon/domain/authdomain"
	"backgammon/utils"
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"regexp"
	"sync"
	"testing"
)

var conf = config.ServerConfig{
	Token: config.TokenConfig{
		TokenLength:  16,
		TokenSymbols: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"},
}

var generator = NewTokenGeneratorFlex(&conf)

func TestGenerateToken(t *testing.T) {
	n := 10000
	tokenRegex := fmt.Sprintf("^[a-zA-Z0-9]{%d}$", conf.Token.TokenLength)

	for i := 0; i < n; i++ {
		token := generator.GenerateToken()
		isValid, err := regexp.MatchString(tokenRegex, string(token))
		if err != nil {
			log.Println(err)
		}
		assert.True(t, isValid)
	}
}

func TestTokenGenerationUniqueness(t *testing.T) {
	n := 10000

	var isUnique bool
	tokens := make([]authdomain.Token, 0, n)

	for i := 0; i < n; i++ {
		token := generator.GenerateToken()

		isUnique = !utils.ContainsToken(&tokens, token)

		assert.True(t, isUnique)
		tokens = append(tokens, token)
	}
}

func TestTokenGenerationUniqueness_Multi(t *testing.T) {
	n := 10000
	m := 6

	var isUnique bool
	tokens := make([]authdomain.Token, 0, n*m)

	mu := sync.Mutex{}

	wg := sync.WaitGroup{}
	wg.Add(m)

	for i := 0; i < m; i++ {
		go func() {
			for j := 0; j < 10000; j++ {
				localToken := generator.GenerateToken()
				mu.Lock()
				tokens = append(tokens, localToken)
				mu.Unlock()
			}
			wg.Done()
		}()
	}

	wg.Wait()

	for i := 0; i < len(tokens); i++ {
		token := tokens[i]
		temp := tokens[i+1:]
		isUnique = !utils.ContainsToken(&temp, token)
		if !isUnique {
			log.Println(token)
		}
		assert.True(t, isUnique)
	}
}

func TestTokenGenerationUniqueness_MultiGenerator(t *testing.T) {
	n := 10000
	m := 6

	var isUnique bool
	tokens := make([]authdomain.Token, 0, n*m)

	mu := sync.Mutex{}

	wg := sync.WaitGroup{}
	wg.Add(m)

	for i := 0; i < m; i++ {
		go func() {
			localGenerator := NewTokenGeneratorFlex(&conf)
			for j := 0; j < 10000; j++ {
				localToken := localGenerator.GenerateToken()
				mu.Lock()
				tokens = append(tokens, localToken)
				mu.Unlock()
			}
			wg.Done()
		}()
	}

	wg.Wait()

	for i := 0; i < len(tokens); i++ {
		token := tokens[i]
		temp := tokens[i+1:]
		isUnique = !utils.ContainsToken(&temp, token)
		if !isUnique {
			log.Println(token)
		}
		assert.True(t, isUnique)
	}
}

func TestTokenGeneratorFlex_Multi(t *testing.T) {
	for i := 0; i < 6; i++ {
		go func() {
			for j := 0; j < 10000; j++ {
				token := generator.GenerateToken()
				assert.Greater(t, len(token), 0)
				assert.True(t, true)
			}
		}()
	}
}
