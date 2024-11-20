package kv

import "time"

type Store interface {
	Get(key string) ([]byte, error)
	Set(key string, value []byte) error
	SetWithTTL(key string, value []byte, ttl time.Duration) error
	SetTTL(key string, ttl int) error
	Delete(key string) error
	Clear() error
	ClearExpire() error
}
