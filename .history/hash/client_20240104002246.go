package hash

import (
	"crypto/sha256"
	"hash"
	"sync"
)

type hashImpl struct {
	hash *hash.Hash
}

var (
	hashInstance *hashImpl
	hashOnce     sync.Once
	errorDb      error
)

func NewHash() (hashInstance hashImpl, err error) {
	hashOnce.Do(func() {
		hs := sha256.New()

		hashInstance = &hashImpl{hs}
	})
	return hashInstance, nil
}
