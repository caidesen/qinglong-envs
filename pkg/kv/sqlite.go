package kv

import (
	"database/sql"
	"time"
)

type SqliteKVStore struct {
	db *sql.DB
}

func NewSqliteKVStore(db *sql.DB) (*SqliteKVStore, error) {
	_, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS kv_store (
            key TEXT PRIMARY KEY,
            value BLOB,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			expired_at DATETIME
        )
    `)
	if err != nil {
		return nil, err
	}
	return &SqliteKVStore{
		db: db,
	}, nil
}

func (store *SqliteKVStore) Get(key string) ([]byte, error) {
	var value []byte
	var expiredAt time.Time
	err := store.db.QueryRow("SELECT value, expired_at FROM kv_store WHERE key = ?", key).Scan(&value, &expiredAt)
	if err == nil && expiredAt.Before(time.Now()) {
		return nil, nil
	}
	return value, err
}

func (store *SqliteKVStore) Set(key string, value []byte) error {
	_, err := store.db.Exec("INSERT OR REPLACE INTO kv_store (key, value) VALUES (?, ?)", key, value)
	return err
}

func (store *SqliteKVStore) SetWithTTL(key string, value []byte, ttl time.Duration) error {
	expiredAt := time.Now().Add(ttl)
	_, err := store.db.Exec("INSERT OR REPLACE INTO kv_store (key, value, expired_at) VALUES (?, ?, datetime(?))", key, value, expiredAt)
	return err
}

func (store *SqliteKVStore) SetTTL(key string, ttl int) error {
	store.db.Exec("UPDATE kv_store SET expired_at = ? WHERE key = ?", ttl, key)
	return nil
}

func (store *SqliteKVStore) Delete(key string) error {
	_, err := store.db.Exec("DELETE FROM kv_store WHERE key = ?", key)
	return err
}

func (store *SqliteKVStore) Clear() error {
	_, err := store.db.Exec("DELETE FROM kv_store")
	return err
}

func (store *SqliteKVStore) ClearExpire() error {
	_, err := store.db.Exec("DELETE FROM kv_store WHERE expired_at < DATETIME('now')")
	return err
}
