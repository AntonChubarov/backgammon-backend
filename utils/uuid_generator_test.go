package utils

import (
	"backgammon/domain/authdomain"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

func TestGenerateUUID(t *testing.T) {
	n := 10000
	var uuid authdomain.UUID
	uuidRegex := "\\b[0-9a-f]{8}\\b-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-\\b[0-9a-f]{12}\\b"

	re := regexp.MustCompile(uuidRegex)

	for i := 0; i < n; i++ {
		uuid = GenerateUUID()
		isValid := re.MatchString(string(uuid))

		assert.True(t, isValid)
	}
}

func TestUUIDGenerationUniqueness(t *testing.T) {
	n := 10000
	var uuid authdomain.UUID
	var isUnique bool
	uuids := make([]string, 0, n)

	for i := 0; i < n; i++ {
		uuid = GenerateUUID()

		isUnique=!Contains(&uuids, string(uuid))

		assert.True(t, isUnique)
		uuids=append(uuids, string(uuid))
	}
}

