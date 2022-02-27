// Package pairStoring contains the heart of the domain model.
package pairStoring

import (
	"errors"
)

var (
	ErrKeyNotFound = errors.New("the key was not found")
)

type Pair struct {
	Key   string
	Value string
}

func NewPair(key, value string) *Pair {
	return &Pair{Key: key, Value: value}
}

type PairRepository interface {
	Set(p *Pair) bool
	Get(key string) (string, error)
	SetAll(m map[string]string)
	GetAll() map[string]string
	Flush()
}
