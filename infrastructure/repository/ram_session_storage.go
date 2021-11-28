package repository

import (
	"backgammon/app/service"
	"backgammon/domain"
	"sync"
)

type SessionStorageRAM struct {
	sessions  map[domain.Token]*domain.SessionData
	uuidIndex map[domain.UUID]*domain.SessionData
	sync.RWMutex
}

func NewSessionStorageRam() *SessionStorageRAM {
	return &SessionStorageRAM{
		sessions:  make(map[domain.Token]*domain.SessionData, 0),
		uuidIndex: make(map[domain.UUID]*domain.SessionData, 0),
		RWMutex:   sync.RWMutex{},
	}
}

func (s *SessionStorageRAM) AddSession(data domain.SessionData) error {
	s.Lock()
	defer s.Unlock()
	_, ok := s.sessions[data.Token]
	if ok {
		return service.ErrorDuplicateSession
	}
	_, ok2 := s.uuidIndex[data.UUID]
	if ok2 {
		return service.ErrorUserMultiSessioning
	}
	copyData := data
	s.sessions[data.Token] = &copyData
	s.uuidIndex[data.UUID] = &copyData
	return nil
}

func (s *SessionStorageRAM) GetSessionByToken(token domain.Token) (domain.SessionData, error) {
	s.Lock()
	defer s.Unlock()
	ses, ok := s.sessions[token]
	if !ok {
		return domain.SessionData{}, service.ErrorInvalidToken
	}
	return *ses, nil
}

func (s *SessionStorageRAM) GetSessionSByUUID(uuid domain.UUID) (domain.SessionData, error) {
	s.Lock()
	defer s.Unlock()
	ses, ok := s.uuidIndex[uuid]
	if !ok {
		return domain.SessionData{}, service.ErrorNoActiveSessions
	}
	return *ses, nil

}

func (s *SessionStorageRAM) DeleteSession(token domain.Token) error {
	s.Lock()
	defer s.Unlock()
	ses, ok := s.sessions[token]
	if !ok {
		return service.ErrorInvalidToken
	}
	delete(s.uuidIndex, ses.UUID)
	delete(s.sessions, token)
	return nil
}

func (s *SessionStorageRAM) UpdateSession(token domain.Token, data domain.SessionData) error {
	s.Lock()
	defer s.Unlock()
	ses, ok := s.sessions[token]
	if !ok {
		return service.ErrorInvalidToken
	}
	ses.ExpiryTime = data.ExpiryTime
	ses.RoomID = data.RoomID
	return nil
}
