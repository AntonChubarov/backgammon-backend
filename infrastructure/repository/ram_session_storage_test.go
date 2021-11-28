package repository

import (
	"backgammon/app/service"
	"backgammon/config"
	"backgammon/domain"
	"backgammon/utils"
	"github.com/stretchr/testify/assert"
	"log"
	"sync"
	"sync/atomic"
	"testing"
)

var serverConfig = config.ServerConfig{
	Token: config.TokenConfig{
		TokenLength:  16,
		TokenSymbols: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789",
	},
}

func TestSessionStorageRAM_AddSession_single(t *testing.T) {

	tokenGenerator := service.NewTokenGeneratorFlex(&serverConfig)
	storage := NewSessionStorageRam()
	token := tokenGenerator.GenerateToken()
	ses := domain.SessionData{
		UUID:       utils.GenerateUUID(),
		Token:      domain.Token(token),
		RoomID:     "",
		ExpiryTime: domain.ExpiryTime{},
	}
	err2 := storage.AddSession(ses)
	assert.Nil(t, err2)
}
func TestSessionStorageRAM_AddSession_single_GetSessionByUUID(t *testing.T) {

	tokenGenerator := service.NewTokenGeneratorFlex(&serverConfig)
	storage := NewSessionStorageRam()
	token := tokenGenerator.GenerateToken()
	ses := domain.SessionData{
		UUID:       utils.GenerateUUID(),
		Token:      domain.Token(token),
		RoomID:     "",
		ExpiryTime: domain.ExpiryTime{},
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
	tokenGenerator := service.NewTokenGeneratorFlex(&serverConfig)
	storage := NewSessionStorageRam()
	token := domain.Token(tokenGenerator.GenerateToken())
	ses, err := storage.GetSessionByToken(token)
	assert.Equal(t, domain.SessionData{}, ses)
	assert.Equal(t, service.ErrorInvalidToken, err)
}

func TestSessionStorageRAM_GetSessionByUUID_absent_uuid(t *testing.T) {
	tokenGenerator := service.NewTokenGeneratorFlex(&serverConfig)
	storage := NewSessionStorageRam()
	uuid := domain.UUID(tokenGenerator.GenerateToken())
	ses, err := storage.GetSessionSByUUID(uuid)
	assert.Equal(t, domain.SessionData{}, ses)
	assert.Equal(t, service.ErrorNoActiveSessions, err)
}

func TestSessionStorageRAM_AddSession_duplicate_session(t *testing.T) {

	tokenGenerator := service.NewTokenGeneratorFlex(&serverConfig)
	storage := NewSessionStorageRam()
	token := tokenGenerator.GenerateToken()
	ses := domain.SessionData{
		UUID:       utils.GenerateUUID(),
		Token:      domain.Token(token),
		RoomID:     "",
		ExpiryTime: domain.ExpiryTime{},
	}

	ses2 := domain.SessionData{
		UUID:       utils.GenerateUUID(),
		Token:      ses.Token,
		RoomID:     "",
		ExpiryTime: domain.ExpiryTime{},
	}

	err2 := storage.AddSession(ses)
	assert.Nil(t, err2)

	err3 := storage.AddSession(ses2)
	assert.Equal(t, service.ErrorDuplicateSession, err3)

}

func TestSessionStorageRAM_AddSession_duplicate_uuid(t *testing.T) {
	tokenGenerator := service.NewTokenGeneratorFlex(&serverConfig)
	storage := NewSessionStorageRam()
	token := domain.Token(tokenGenerator.GenerateToken())
	token2 := domain.Token(tokenGenerator.GenerateToken())
	ses := domain.SessionData{
		UUID:       utils.GenerateUUID(),
		Token:      token,
		RoomID:     "",
		ExpiryTime: domain.ExpiryTime{},
	}

	ses2 := domain.SessionData{
		UUID:       ses.UUID,
		Token:      token2,
		RoomID:     "",
		ExpiryTime: domain.ExpiryTime{},
	}

	err2 := storage.AddSession(ses)
	assert.Nil(t, err2)

	err3 := storage.AddSession(ses2)
	assert.Equal(t, service.ErrorUserMultiSessioning, err3)

}

func TestSessionStorageRAM_DeleteSession(t *testing.T) {
	tokenGenerator := service.NewTokenGeneratorFlex(&serverConfig)
	storage := NewSessionStorageRam()
	token := domain.Token(tokenGenerator.GenerateToken())
	ses := domain.SessionData{
		UUID:       utils.GenerateUUID(),
		Token:      token,
		RoomID:     "",
		ExpiryTime: domain.ExpiryTime{},
	}
	err := storage.AddSession(ses)
	assert.Nil(t, err)

	err2 := storage.DeleteSession(ses.Token)
	assert.Nil(t, err2)

	ses2, err3 := storage.GetSessionByToken(ses.Token)
	assert.Equal(t, domain.SessionData{}, ses2)
	assert.Equal(t, service.ErrorInvalidToken, err3)

	ses3, err4 := storage.GetSessionSByUUID(ses.UUID)
	assert.Equal(t, domain.SessionData{}, ses3)
	assert.Equal(t, service.ErrorNoActiveSessions, err4)
}

func TestSessionStorageRAM_UpdateSession(t *testing.T) {
	tokenGenerator := service.NewTokenGeneratorFlex(&serverConfig)
	storage := NewSessionStorageRam()
	token := domain.Token(tokenGenerator.GenerateToken())
	token2 := domain.Token(tokenGenerator.GenerateToken())
	ses := domain.SessionData{
		UUID:       utils.GenerateUUID(),
		Token:      token,
		RoomID:     domain.RoomID(tokenGenerator.GenerateToken()),
		ExpiryTime: domain.ExpiryTime{},
	}

	ses2 := domain.SessionData{
		UUID:       ses.UUID,
		Token:      token2,
		RoomID:     domain.RoomID(tokenGenerator.GenerateToken()),
		ExpiryTime: domain.ExpiryTime{},
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
	wg := sync.WaitGroup{}

	sl1 := fillSessionSlice(10000)
	sl2 := fillSessionSlice(10000)
	sl3 := fillSessionSlice(10000)
	sl4 := fillSessionSlice(10000)
	sl5 := fillSessionSlice(10000)
	fnAdd := func(s domain.SessionStorage, sl []domain.SessionData) {
		for i := range sl {
			s.AddSession(sl[i])
		}
		wg.Done()
	}

	storage := NewSessionStorageRam()
	wg.Add(4)
	fnAdd(storage, sl1) //to read by Token
	fnAdd(storage, sl2) //to read by UUID
	fnAdd(storage, sl3) //To delete
	fnAdd(storage, sl4) //to update from read by Token

	fnUpd := func(slOrg []domain.SessionData, slUpd []domain.SessionData) {
		for i := range slOrg {
			err := storage.UpdateSession(sl4[i].Token, sl2[i])
			assert.Nil(t, err)
			ses, err2 := storage.GetSessionByToken(slOrg[i].Token)
			assert.Nil(t, err2)
			assert.Equal(t, slOrg[i].Token, ses.Token)
			assert.Equal(t, slUpd[i].RoomID, ses.RoomID)
			assert.Equal(t, slUpd[i].ExpiryTime, ses.ExpiryTime)
		}
		wg.Done()
	}

	fnDel := func(sl []domain.SessionData) {
		for i := range sl {
			err := storage.DeleteSession(sl[i].Token)
			assert.Nil(t, err)
		}
		wg.Done()
	}

	fnRdToken := func(sl []domain.SessionData) {
		for i := range sl {
			ses, err := storage.GetSessionByToken(sl[i].Token)
			assert.Nil(t, err)
			assert.Equal(t, sl[i], ses)
		}
		wg.Done()
	}

	fnRdUuid := func(sl []domain.SessionData) {
		for i := range sl {
			ses, err := storage.GetSessionSByUUID(sl[i].UUID)
			assert.Nil(t, err)
			assert.Equal(t, sl[i], ses)
		}
		wg.Done()
	}

	wg.Add(5)
	go fnRdToken(sl1)
	go fnRdUuid(sl2)
	go fnDel(sl3)
	go fnUpd(sl4, sl2)
	go fnAdd(storage, sl5)

	wg.Wait()

	assert.True(t, true)

}

func TestSessionStorageRAM_MultiRoutine_Interleave(t *testing.T) {
	wg := sync.WaitGroup{}
	var readsCounter int64
	maxInterval := 100
	sl1 := fillSessionSlice(10000)
	sl2 := fillSessionSlice(10000)
	sl3 := fillSessionSlice(10000)
	sl4 := fillSessionSlice(10000)
	sl5 := fillSessionSlice(10000)
	fnAdd := func(s domain.SessionStorage, sl []domain.SessionData) {
		for i := range sl {
			s.AddSession(sl[i])
			tmp := atomic.LoadInt64(&readsCounter)
			atomic.StoreInt64(&readsCounter, 0)
			assert.Less(t, tmp, int64(maxInterval))
		}
		wg.Done()
	}

	storage := NewSessionStorageRam()
	wg.Add(4)
	fnAdd(storage, sl1) //to read by Token
	fnAdd(storage, sl2) //to read by UUID
	fnAdd(storage, sl3) //To delete
	fnAdd(storage, sl4) //to update from read by Token
	//-----

	fnUpd := func(slOrg []domain.SessionData, slUpd []domain.SessionData) {
		for i := range slOrg {
			err := storage.UpdateSession(sl4[i].Token, sl2[i])
			tmp := atomic.LoadInt64(&readsCounter)
			atomic.StoreInt64(&readsCounter, 0)
			assert.Less(t, tmp, int64(maxInterval))
			assert.Nil(t, err)
			ses, err2 := storage.GetSessionByToken(slOrg[i].Token)
			assert.Nil(t, err2)
			assert.Equal(t, slOrg[i].Token, ses.Token)
			assert.Equal(t, slUpd[i].RoomID, ses.RoomID)
			assert.Equal(t, slUpd[i].ExpiryTime, ses.ExpiryTime)
		}
		wg.Done()
	}

	fnDel := func(sl []domain.SessionData) {
		for i := range sl {
			err := storage.DeleteSession(sl[i].Token)
			tmp := atomic.LoadInt64(&readsCounter)
			atomic.StoreInt64(&readsCounter, 0)
			assert.Less(t, tmp, int64(maxInterval))
			assert.Nil(t, err)
		}
		wg.Done()
	}

	fnRdToken := func(sl []domain.SessionData) {
		for i := range sl {
			atomic.AddInt64(&readsCounter, 1)
			ses, err := storage.GetSessionByToken(sl[i].Token)
			assert.Nil(t, err)
			assert.Equal(t, sl[i], ses)
		}
		wg.Done()
	}

	fnRdUuid := func(sl []domain.SessionData) {
		for i := range sl {
			ses, err := storage.GetSessionSByUUID(sl[i].UUID)
			atomic.AddInt64(&readsCounter, 1)
			assert.Nil(t, err)
			assert.Equal(t, sl[i], ses)
		}
		wg.Done()
	}

	wg.Add(5)
	go fnRdToken(sl1)
	go fnRdUuid(sl2)
	go fnDel(sl3)
	go fnUpd(sl4, sl2)
	go fnAdd(storage, sl5)

	wg.Wait()

	assert.True(t, true)
	log.Println(readsCounter)

}

func fillSessionSlice(count int) []domain.SessionData {
	tokenGenerator := service.NewTokenGeneratorFlex(&serverConfig)
	sl := make([]domain.SessionData, count, count)
	for i := 0; i < count; i++ {
		ses := domain.SessionData{
			UUID:       utils.GenerateUUID(),
			Token:      domain.Token(tokenGenerator.GenerateToken()),
			RoomID:     domain.RoomID(tokenGenerator.GenerateToken()),
			ExpiryTime: domain.ExpiryTime{},
		}
		sl[i] = ses
	}
	return sl
}
