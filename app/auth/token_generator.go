package auth

import (
	"backgammon/config"
	"backgammon/utils"

)

type TokenGeneratorFlex struct {
	TokenLength int
	AllowedSymbols string
}

func NewTokenGeneratorFlex(c *config.ServerConfig ) TokenGenerator {
	return &TokenGeneratorFlex{TokenLength: c.Token.TokenLength, AllowedSymbols: c.Token.TokenSymbols}
}

func (tgf *TokenGeneratorFlex) GenerateToken() (token string) {
	token = ""
	var i int
	for len(token) < tgf.TokenLength {
		i = utils.RandomInt(len(tgf.AllowedSymbols))
		token += string(tgf.AllowedSymbols[i])
	}
	return
}

