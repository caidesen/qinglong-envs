APP_NAME = qinglong-envs
BUILD_DIR = $(PWD)/build

build: clean
	CGO_ENABLED=1 go build -ldflags="-w -s" -o $(BUILD_DIR)/$(APP_NAME)_linux_amd64 cmd/server/main.go
	#CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags="-w -s" -o $(BUILD_DIR)/$(APP_NAME)_linux_arm64 main.go

mkdir:
	mkdir -p build

clean:
	rm -rf ./build
