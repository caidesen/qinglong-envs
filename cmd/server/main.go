package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log/slog"
	"os"
	"path"
	"qinglong-envs/internal/handlers"
	"qinglong-envs/internal/services"
	"qinglong-envs/pkg/api"
	"qinglong-envs/pkg/kv"
	"qinglong-envs/pkg/middleware"
	"qinglong-envs/pkg/router"
)

func checkPath(p string) string {
	if path.IsAbs(p) {
		return p
	}
	dir, _ := os.Getwd()
	return path.Join(dir, p)
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
	kvStore := InitKVStore(InitDB("./.tmp/kv.db"))
	// service 初始化
	tokenService := services.NewTokenService(kvStore)
	userService := services.NewUserService(db, tokenService)
	// 注册路由
	r := router.NewRouter()
	r.Use(middleware.Recovery)
	r.GroupWithPrefix("/api", func(ar router.Router) {
		ar.Group(handlers.NewUserHandler(userService).Register)
	})
	api.StartHttpServer(r)
}
