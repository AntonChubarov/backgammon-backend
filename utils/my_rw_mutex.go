package utils

import (
	"sync"
)

type MyRWMutex struct {
	encapsulatedRWMutex sync.RWMutex
	encapsulatedMutex   sync.Mutex
	readLocksCounter    int64
}

//func (m *MyRWMutex) GetNumberRLocks() int {
//
//}
