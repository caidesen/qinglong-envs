package env

import (
	"fmt"
	"github.com/joho/godotenv"
	"log/slog"
	"os"
	"path/filepath"
)

// Load loads the environment variables from the .env file.
func Load(envFile string) {
	path, err := dir(envFile)
	if err != nil {
		slog.Warn("load env file: %s", err)
	}
	err = godotenv.Load(path)
	if err != nil {
		slog.Warn("load env file: %s", err)
	}
}

// dir returns the absolute path of the given environment file (envFile) in the Go module's
// root directory. It searches for the 'go.mod' file from the current working directory upwards
// and appends the envFile to the directory containing 'go.mod'.
// It panics if it fails to find the 'go.mod' file.
func dir(envFile string) (string, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	level := 0
	for {
		goModPath := filepath.Join(currentDir, "go.mod")
		if _, err := os.Stat(goModPath); err == nil {
			break
		}

		parent := filepath.Dir(currentDir)
		if parent == currentDir && level > 4 {
			return "", fmt.Errorf("go.mod not found")
		}
		currentDir = parent
		level++
	}

	return filepath.Join(currentDir, envFile), nil
}

func init() {
	Load(".env")
}
