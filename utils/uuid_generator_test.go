package utils

import (
	"github.com/stretchr/testify/assert"
	"log"
	"regexp"
	"testing"
)

func TestGenerateUUID(t *testing.T) {
	n := 1000
	var uuid string
	uuidRegex := "\\b[0-9a-f]{8}\\b-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-\\b[0-9a-f]{12}\\b"

	for i := 0; i < n; i++ {
		uuid = GenerateUUID()
		isValid, err := regexp.MatchString(uuidRegex, uuid)
		if err != nil {
			log.Println(err)
		}
		assert.True(t, isValid)
	}
}

func TestUUIDGenerationUniqueness(t *testing.T) {
	n := 10000
	var uuid string
	var isUnique bool
	uuids := make([]string, 0, n)

	for i := 0; i < n; i++ {
		uuid = GenerateUUID()

		isUnique=!Contains(&uuids, uuid)

		assert.True(t, isUnique)
		uuids=append(uuids, uuid)
	}
}

