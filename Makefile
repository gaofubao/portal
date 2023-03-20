.PHONY: build run clean help

GO_BIN="go"
APP_NAME="portal"
MODEL_DDL_PATH="internal/model"

## build: Build source code for host platform.
.PHONY: build
build:
	@${GO_BIN} build -o "${APP_NAME}.exe" portal.go

## linux: Build source code for linux platform.
.PHONY: linux
linux:
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 ${GO_BIN} build -o "${APP_NAME}-linux.exe" portal.go

## run: Run the project.
.PHONY: run
run:
	@${GO_BIN} run portal.go

## clean: Clean binary file.
.PHONY: clean
clean:
	@rm -rf *.exe

## api: generate grpc code.
.PHONY: api
api:
	@goctl api go --api api/portal.api --dir . --style=go_zero

## model: generate model code
.PHONY: model
model:
	@cd ${MODEL_DDL_PATH} && goctl model mysql ddl --src sql/*.sql --dir . --cache

## help: Help.
.PHONY: help
help:
	@echo "make build 构建当前平台的二进制文件"
	@echo "make linux 构建 Linux 平台的二进制文件"
	@echo "make run 运行"
	@echo "make clean 清除构建的二进制文件"
	@echo "make api 生成 api 代码"
	@echo "make model 生成 model 代码"
