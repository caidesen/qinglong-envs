package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"qinglong-envs/internal/db/queries"
	"qinglong-envs/internal/server/api/v1"
	"qinglong-envs/pkg/kv"
	"qinglong-envs/pkg/middleware"
	"qinglong-envs/pkg/router"
	"strconv"
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

func LoadListeningFormEnv() string {
	port := 3000
	host := "127.0.0.1"
	portEnv := os.Getenv("QINGLONG_ENVS_SERVER_PORT")
	hostEnv := os.Getenv("QINGLONG_ENVS_SERVER_HOST")
	// to int
	portEnvInt, _ := strconv.Atoi(portEnv)
	if portEnvInt != 0 {
		port = portEnvInt
	}
	if hostEnv != "" {
		host = hostEnv
	}
	return fmt.Sprintf("%s:%d", host, port)
}

func StartHttpServer(h http.Handler) {
	addr := LoadListeningFormEnv()
	server := &http.Server{
		Addr:    addr,
		Handler: h,
	}
	slog.Info(fmt.Sprintf("listening on %s", addr))
	log.Fatal(server.ListenAndServe())
}

func main() {
	// 基础组件初始化
	db := InitDB("./.tmp/data.db")
	q := queries.New(db)
	// 注册路由
	r := router.New(http.NewServeMux())
	r.Use(middleware.Recovery)
	r.Mount("/api/v1").Route(v1.NewHandlers(q).Routes)
	StartHttpServer(r)
}
