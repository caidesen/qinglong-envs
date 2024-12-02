package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	apiv1 "qinglong-envs/internal/api/v1"
	"qinglong-envs/internal/db/queries"
	"qinglong-envs/pkg/httpapi"
	"qinglong-envs/pkg/kv"
	"qinglong-envs/pkg/middleware"
)

func checkPath(p string) string {
	if filepath.IsAbs(p) {
		return p
	}
	dir, _ := os.Getwd()
	return filepath.Join(dir, p)
}

func InitDB(dbPath string) *sql.DB {
	dbPath = checkPath(dbPath)
	formatPath := fmt.Sprintf("%s?_journal=WAL&_vacuum=incremental", dbPath)
	dB, err := sql.Open("sqlite3", formatPath)
	if err != nil {
		slog.Error(fmt.Sprintf("open db error: %v", err))
		os.Exit(1)
	}
	err = dB.Ping()
	if err != nil {
		slog.Error(fmt.Sprintf("ping db error: %v", err))
		os.Exit(1)
	}
	return dB
}

func InitKVStore(db *sql.DB) kv.Store {
	store, err := kv.NewSqliteKVStore(db)
	if err != nil {
		slog.Error(fmt.Sprintf("open db error: %v", err))
		os.Exit(1)
	}
	return store
}

func main() {
	// 基础组件初始化
	db := InitDB("./.tmp/data.db")
	q := queries.New(db)
	v1 := apiv1.New(q)
	// 注册路由
	router := httpapi.NewRouter()
	router.Use(middleware.Recovery)

	router.Group(func(r httpapi.Router) {
		v1.Register(r)
	})

	mux := http.NewServeMux()
	mux.Handle("/api/v1/", http.StripPrefix("/api/v1/", router))
	//router.Handle("/api/v1/", http.StripPrefix("/api/v1/", router))

	httpapi.StartHttpServer(mux)
}
