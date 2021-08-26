package utils

import (
	"backgammon/domain/authdomain"
	"crypto/rand"
	"fmt"
	"log"
)

func GenerateUUID() authdomain.UUID {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return authdomain.UUID(uuid)
}
