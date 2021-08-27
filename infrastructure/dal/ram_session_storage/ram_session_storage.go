package ram_session_storage

import (
	"backgammon/app/auth"
	"backgammon/domain/authdomain"
	"sync"
)

type SessionStorageRAM struct {
	sessions map[authdomain.Token]*authdomain.SessionData
	uuidIndex map[authdomain.UUID]*authdomain.SessionData
	sync.RWMutex
}

func NewSessionStorageRam() *SessionStorageRAM {
	return &SessionStorageRAM{
		sessions:  make(map[authdomain.Token]*authdomain.SessionData, 0),
		uuidIndex: make(map[authdomain.UUID]*authdomain.SessionData,0),
		RWMutex:   sync.RWMutex{},
	}
}

func (s *SessionStorageRAM) AddSession(data authdomain.SessionData) error {
	s.Lock()
	defer s.Unlock()
	_, ok:=s.sessions[data.Token]
	if ok {return auth.ErrorDuplicateSession}
	_, ok2:=s.uuidIndex[data.UUID]
	if ok2 {return auth.ErrorUserMultiSessioning}
	copyData:=data
	s.sessions[data.Token]=&copyData
	s.uuidIndex[data.UUID]=&copyData
	return nil
}

func (s *SessionStorageRAM) GetSessionByToken(token authdomain.Token) (authdomain.SessionData, error) {
	s.Lock()
	defer s.Unlock()
	ses, ok:=s.sessions[token]
	if !ok {return authdomain.SessionData{}, auth.ErrorInvalidToken}
	return *ses, nil
}

func (s *SessionStorageRAM) GetSessionSByUUID(uuid authdomain.UUID) (authdomain.SessionData, error) {
	s.Lock()
	defer s.Unlock()
	ses, ok:=s.uuidIndex[uuid]
	if !ok { return authdomain.SessionData{}, auth.ErrorNoActiveSessions}
	return *ses, nil

}

func (s *SessionStorageRAM) DeleteSession(token authdomain.Token) error {
	s.Lock()
	defer s.Unlock()
	ses, ok:=s.sessions[token]
	if !ok {return auth.ErrorInvalidToken}
	delete(s.uuidIndex, ses.UUID)
	delete(s.sessions, token)
	return nil
}

func (s *SessionStorageRAM) UpdateSession(token authdomain.Token, data authdomain.SessionData) error {
	s.Lock()
	defer s.Unlock()
	ses, ok:=s.sessions[token]
	if !ok  {return auth.ErrorInvalidToken}
	ses.ExpiryTime=data.ExpiryTime
	ses.RoomID=data.RoomID
	return nil
}






