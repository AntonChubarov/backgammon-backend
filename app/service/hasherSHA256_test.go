package service

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

type testCase struct {
	Password string
	Hash     string
}

func TestHasherSHA256_Single(t *testing.T) {
	cases := []testCase{
		{
			Password: "vasa2004",
			Hash:     "2b6eb6977d6650d3ab292344b35faca55feccad89409056de06b8831ab01c101",
		},
		{
			Password: "",
			Hash:     "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
		},
		{
			Password: "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
			Hash:     "cd372fb85148700fa88095e3492d3f9f5beb43e555e5ff26d95f5a6adc36f8e6",
		},
	}

	hasher := NewHasherSHA256()

	for i := range cases {
		expected := cases[i].Hash
		actual, err := hasher.HashString(cases[i].Password)
		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	}
}

func TestHasherSHA256_Sequential(t *testing.T) {
	n := 100000
	var password string
	hasher := NewHasherSHA256()
	cases := make([]testCase, 0, n)
	for i := 0; i < n; i++ {
		password = gofakeit.Password(true, true, true, false, false, 10)
		hash, err := hasher.HashString(password)
		assert.Nil(t, err)
		cases = append(cases, testCase{Password: password, Hash: hash})
	}

	for i := range cases {
		expected := cases[i].Hash
		actual, err := hasher.HashString(cases[i].Password)
		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	}
}

func TestHasherSHA256_Massive(t *testing.T) {
	n := 1000000
	var password string
	hasher := NewHasherSHA256()
	cases := make([]testCase, 0, n)
	for i := 0; i < n; i++ {
		password = gofakeit.Password(true, true, true, false, false, 10)
		hash, err := hasher.HashString(password)
		assert.Nil(t, err)
		cases = append(cases, testCase{Password: password, Hash: hash})
	}

	casesPackage := [][]testCase{
		cases[0:99999],
		cases[100000:199999],
		cases[200000:299999],
		cases[300000:399999],
		cases[400000:499999],
		cases[500000:599999],
		cases[600000:699999],
		cases[700000:799999],
		cases[800000:899999],
		cases[900000:999999],
	}

	wg := sync.WaitGroup{}
	wg.Add(len(casesPackage))

	for j := 0; j < len(casesPackage); j++ {
		go func(partialCases []testCase) {
			for i := range partialCases {
				expected := cases[i].Hash
				actual, err := hasher.HashString(cases[i].Password)
				assert.Nil(t, err)
				assert.Equal(t, expected, actual)
			}
			wg.Done()
		}(casesPackage[j])
	}

	wg.Wait()
}
