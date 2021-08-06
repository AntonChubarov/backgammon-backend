package app

import (
	"crypto/sha256"
	"fmt"
	"log"
)

func HashPassword(password string) (string, error) {
	h := sha256.New()
	_, err := h.Write([]byte(password))
	if err != nil {
		log.Println("In app.HashPassword", err)
		return "", err
	}

	hs := fmt.Sprintf("%x", h.Sum(nil))
	return hs, nil
}
