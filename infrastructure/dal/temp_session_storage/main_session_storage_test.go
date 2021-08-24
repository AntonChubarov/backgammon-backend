package temp_session_storage

import (
	"backgammon/app/auth"
	"backgammon/config"
	"backgammon/utils"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var storage = NewMainSessionStorage()
var tokenGenerator = auth.NewTokenGeneratorFlex(
	&config.ServerConfig{
		Token: config.TokenConfig{
			TokenLength: 16,
			TokenSymbols: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"},
	},
	)

func TestAddNewUser(t *testing.T) {
	var token string
	var data auth.UserSessionData


	for i := 0; i < 8; i++ {
		//go func() {
			for i := 0; i < 10; i++ {
				data = auth.UserSessionData{
					Token:      tokenGenerator.GenerateToken(),
					ExpiryTime: time.Now().UTC().Add(30 * time.Second),
					UserUUID:   string(utils.GenerateUUID()),
					WebSocket:  &websocket.Conn{},
				}

				storage.AddNewUser(&data)
				token, _ = storage.GetTokenByUUID(data.UserUUID)
				fmt.Println(token)
				assert.Equal(t, data.Token, token)
			}
		//}()
	}

	time.Sleep(1 * time.Second)
}