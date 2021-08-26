package auth

import (
	"crypto/sha256"
	"fmt"
	"hash"
	"sync"
)

type HasherSHA256 struct {
	hasher hash.Hash
	mutex  sync.Mutex
}

func NewHasherSHA256() *HasherSHA256 {
	return &HasherSHA256{hasher: sha256.New()}
}

func (h *HasherSHA256) HashString(password string) (string, error) {
	h.mutex.Lock()
	h.hasher.Reset()
	_, err := h.hasher.Write([]byte(password))
	if err != nil {
		//log.Println("In app.HashPassword", err)
		return "", err
	}
	hs := fmt.Sprintf("%x", h.hasher.Sum(nil))
	h.mutex.Unlock()
	return hs, nil
}
