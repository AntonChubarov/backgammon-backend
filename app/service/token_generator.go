package service

import (
	"backgammon/config"
	"backgammon/domain"
	"backgammon/utils"
)

type TokenGeneratorFlex struct {
	TokenLength    int
	AllowedSymbols string
}

func NewTokenGeneratorFlex(c *config.ServerConfig) TokenGenerator {
	return &TokenGeneratorFlex{TokenLength: c.Token.TokenLength, AllowedSymbols: c.Token.TokenSymbols}
}

func (tgf *TokenGeneratorFlex) generateRandomString() (token string) {
	token = ""
	var i int
	for len(token) < tgf.TokenLength {
		i = utils.RandomInt(len(tgf.AllowedSymbols))
		token += string(tgf.AllowedSymbols[i])
	}
	return
}

func (tgf *TokenGeneratorFlex) GenerateToken() domain.Token {
	return domain.Token(tgf.generateRandomString())
}
