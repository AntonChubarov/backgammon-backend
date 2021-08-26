package ram_session_storage

import (
	"backgammon/app/auth"
	"backgammon/config"
	"backgammon/domain/authdomain"
	"backgammon/utils"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

var serverConfig = config.ServerConfig{
	Token: config.TokenConfig{
		TokenLength:  16,
		TokenSymbols: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789",
	},
}

func TestSessionStorageRAM_AddSession_single(t *testing.T) {

	tokenGenerator := auth.NewTokenGeneratorFlex(&serverConfig)
	storage := NewSessionStorageRam()
	token := tokenGenerator.GenerateToken()
	ses := authdomain.SessionData{
		UUID:       utils.GenerateUUID(),
		Token:      authdomain.Token(token),
		RoomID:     "",
		ExpiryTime: authdomain.ExpiryTime{},
	}
	err2 := storage.AddSession(ses)
	assert.Nil(t, err2)
}
func TestSessionStorageRAM_AddSession_single_GetSessionByUUID(t *testing.T) {

	tokenGenerator := auth.NewTokenGeneratorFlex(&serverConfig)
	storage := NewSessionStorageRam()
	token := tokenGenerator.GenerateToken()
	ses := authdomain.SessionData{
		UUID:       utils.GenerateUUID(),
		Token:      authdomain.Token(token),
		RoomID:     "",
		ExpiryTime: authdomain.ExpiryTime{},
	}
	err2 := storage.AddSession(ses)
	assert.Nil(t, err2)

	getSes1, err3 := storage.GetSessionByToken(ses.Token)
	assert.Nil(t, err3)
	assert.Equal(t, ses, getSes1)

	getSes2, err4 := storage.GetSessionSByUUID(ses.UUID)
	assert.Nil(t, err4)
	assert.Equal(t, ses, getSes2)

}

func TestSessionStorageRAM_GetSessionByToken_absent_token(t *testing.T) {
	tokenGenerator := auth.NewTokenGeneratorFlex(&serverConfig)
	storage := NewSessionStorageRam()
	token := authdomain.Token(tokenGenerator.GenerateToken())
	ses, err := storage.GetSessionByToken(token)
	assert.Equal(t, authdomain.SessionData{}, ses)
	assert.Equal(t, auth.ErrorInvalidToken, err)
}

func TestSessionStorageRAM_GetSessionByUUID_absent_uuid(t *testing.T) {
	tokenGenerator := auth.NewTokenGeneratorFlex(&serverConfig)
	storage := NewSessionStorageRam()
	uuid := authdomain.UUID(tokenGenerator.GenerateToken())
	ses, err := storage.GetSessionSByUUID(uuid)
	assert.Equal(t, authdomain.SessionData{}, ses)
	assert.Equal(t, auth.ErrorNoActiveSessions, err)
}

func TestSessionStorageRAM_AddSession_duplicate_session(t *testing.T) {

	tokenGenerator := auth.NewTokenGeneratorFlex(&serverConfig)
	storage := NewSessionStorageRam()
	token := tokenGenerator.GenerateToken()
	ses := authdomain.SessionData{
		UUID:       utils.GenerateUUID(),
		Token:      authdomain.Token(token),
		RoomID:     "",
		ExpiryTime: authdomain.ExpiryTime{},
	}

	ses2 := authdomain.SessionData{
		UUID:       utils.GenerateUUID(),
		Token:      ses.Token,
		RoomID:     "",
		ExpiryTime: authdomain.ExpiryTime{},
	}

	err2 := storage.AddSession(ses)
	assert.Nil(t, err2)

	err3 := storage.AddSession(ses2)
	assert.Equal(t, auth.ErrorDuplicateSession, err3)

}

func TestSessionStorageRAM_AddSession_duplicate_uuid(t *testing.T) {
	tokenGenerator := auth.NewTokenGeneratorFlex(&serverConfig)
	storage := NewSessionStorageRam()
	token := authdomain.Token(tokenGenerator.GenerateToken())
	token2 := authdomain.Token(tokenGenerator.GenerateToken())
	ses := authdomain.SessionData{
		UUID:       utils.GenerateUUID(),
		Token:      token,
		RoomID:     "",
		ExpiryTime: authdomain.ExpiryTime{},
	}

	ses2 := authdomain.SessionData{
		UUID:       ses.UUID,
		Token:      token2,
		RoomID:     "",
		ExpiryTime: authdomain.ExpiryTime{},
	}

	err2 := storage.AddSession(ses)
	assert.Nil(t, err2)

	err3 := storage.AddSession(ses2)
	assert.Equal(t, auth.ErrorUserMultiSessioning, err3)

}

func TestSessionStorageRAM_DeleteSession(t *testing.T) {
	tokenGenerator := auth.NewTokenGeneratorFlex(&serverConfig)
	storage := NewSessionStorageRam()
	token := authdomain.Token(tokenGenerator.GenerateToken())
	ses := authdomain.SessionData{
		UUID:       utils.GenerateUUID(),
		Token:      token,
		RoomID:     "",
		ExpiryTime: authdomain.ExpiryTime{},
	}
	err := storage.AddSession(ses)
	assert.Nil(t, err)

	err2 := storage.DeleteSession(ses.Token)
	assert.Nil(t, err2)

	ses2, err3 := storage.GetSessionByToken(ses.Token)
	assert.Equal(t, authdomain.SessionData{}, ses2)
	assert.Equal(t, auth.ErrorInvalidToken, err3)

	ses3, err4 := storage.GetSessionSByUUID(ses.UUID)
	assert.Equal(t, authdomain.SessionData{}, ses3)
	assert.Equal(t, auth.ErrorNoActiveSessions, err4)
}

func TestSessionStorageRAM_UpdateSession(t *testing.T) {
	tokenGenerator := auth.NewTokenGeneratorFlex(&serverConfig)
	storage := NewSessionStorageRam()
	token := authdomain.Token(tokenGenerator.GenerateToken())
	token2 := authdomain.Token(tokenGenerator.GenerateToken())
	ses := authdomain.SessionData{
		UUID:       utils.GenerateUUID(),
		Token:      token,
		RoomID:     authdomain.RoomID(tokenGenerator.GenerateToken()),
		ExpiryTime: authdomain.ExpiryTime{},
	}

	ses2 := authdomain.SessionData{
		UUID:       ses.UUID,
		Token:      token2,
		RoomID:     authdomain.RoomID(tokenGenerator.GenerateToken()),
		ExpiryTime: authdomain.ExpiryTime{},
	}

	storage.AddSession(ses)
	storage.AddSession(ses2)
	err := storage.UpdateSession(ses.Token, ses2)
	assert.Nil(t, err)

	sesUpd, err2 := storage.GetSessionSByUUID(ses.UUID)
	assert.Nil(t, err2)
	assert.Equal(t, ses.Token, sesUpd.Token)
	assert.Equal(t, ses2.RoomID, sesUpd.RoomID)
	assert.Equal(t, ses2.ExpiryTime, sesUpd.ExpiryTime)
}

func TestSessionStorageRAM_MultiRoutine(t *testing.T) {
	wg:=sync.WaitGroup{}
	wg.Add(5)
	sl1 := fillSessionSlice(10000)
	sl2 := fillSessionSlice(10000)
	sl3 := fillSessionSlice(10000)
	sl4 := fillSessionSlice(10000)
	sl5 := fillSessionSlice(10000)
	fnAdd:=func(s authdomain.SessionStorage, sl []authdomain.SessionData) {
		for i:=range sl {
			s.AddSession(sl[i])
		}
		wg.Done()
	}

	storage := NewSessionStorageRam()
	fnAdd(storage, sl1) //to read by Token
	fnAdd(storage, sl2) //to read by UUID
	fnAdd(storage, sl3) //To delete
	fnAdd(storage, sl4) //to update from read by Token

	fnUpd:= func(slOrg []authdomain.SessionData, slUpd []authdomain.SessionData) {
		for i:= range slOrg {
			err:=storage.UpdateSession(sl4[i].Token, sl2[i])
			assert.Nil(t, err)
			ses, err2:=storage.GetSessionByToken(slOrg[i].Token)
			assert.Nil(t, err2)
			assert.Equal(t, slOrg[i].Token, ses.Token)
			assert.Equal(t, slUpd[i].RoomID, ses.RoomID)
			assert.Equal(t, slUpd[i].ExpiryTime, ses.ExpiryTime)
		}
		wg.Done()
	}

	fnDel:=func (sl []authdomain.SessionData) {
		for i:=range sl{
			err:=storage.DeleteSession(sl[i].Token)
			assert.Nil(t, err)
		}
		wg.Done()
	}

	fnRdToken:=func (sl []authdomain.SessionData) {
		for i:=range sl{
			ses, err:=storage.GetSessionByToken(sl[i].Token)
			assert.Nil(t, err)
			assert.Equal(t, sl[i], ses)
		}
		wg.Done()
	}

	fnRdUuid:=func (sl []authdomain.SessionData) {
		for i:=range sl{
			ses, err:=storage.GetSessionSByUUID(sl[i].UUID)
			assert.Nil(t, err)
			assert.Equal(t, sl[i], ses)
		}
		wg.Done()
	}

	wg.Add(5)
	 fnRdToken(sl1)
	 fnRdUuid(sl2)
	 go fnDel(sl3)
	 fnUpd(sl4, sl2)
	 fnAdd(storage, sl5)

	wg.Wait()

	assert.True(t, true)


}

func fillSessionSlice(count int) []authdomain.SessionData {
	tokenGenerator := auth.NewTokenGeneratorFlex(&serverConfig)
	sl:=make([]authdomain.SessionData, count, count)
	for i:=0; i<count; i++ {
		ses := authdomain.SessionData{
			UUID:       utils.GenerateUUID(),
			Token:      authdomain.Token(tokenGenerator.GenerateToken()),
			RoomID:     authdomain.RoomID(tokenGenerator.GenerateToken()),
			ExpiryTime: authdomain.ExpiryTime{},
		}
		sl[i]=ses
	}
	return sl
}