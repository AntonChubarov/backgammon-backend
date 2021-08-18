package auth

import (
	"crypto/sha256"
	"fmt"
	"hash"
	"log"
)

type HasherSHA256 struct {
	 hasher hash.Hash
}

func NewHasherSHA256() *HasherSHA256 {
	return &HasherSHA256{hasher: sha256.New()}
}

func (h *HasherSHA256) HashString(password string) (string, error) {
	h.hasher.Reset()
	_, err := h.hasher.Write([]byte(password))
	if err != nil {
		log.Println("In app.HashPassword", err)
		return "", err
	}
	hs := fmt.Sprintf("%x", h.hasher.Sum(nil))
	return hs, nil
}

