package hash

import (
	"context"
	"crypto/sha256"
	"hash"
	"sync"
)

type hashImpl struct {
	Hash *hash.Hash
}

var (
	hashInstance *hashImpl
	hashOnce     sync.Once
	errorHash    error
)

func NewHash(ctx context.Context) *hashImpl {
	hashOnce.Do(func() {
		hash := sha256.New()

		hashInstance = &hashImpl{&hash}
	})

	return hashInstance
}
