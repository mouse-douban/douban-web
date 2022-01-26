#!/bin/bash

# jenkins 的持续集成
# 因为使用来热重载，将不再构建 windows 版本

echo "Local dir $PWD"

echo "Setting goproxy..."

go env -w GOPROXY=https://goproxy.cn,direct
go env -w GO111MODULE=on
go mod tidy

echo "Testing ..."

go test -v test/*.go

echo "Building package ..."

echo "Build linux..."
go env -w GOARCH=amd64
go env -w GOOS=linux
go build -o target/linux/bin/gin ./cmd/main.go

echo "Build macos..."
go env -w GOOS=darwin
go build -o target/macos_amd64/bin/gin ./cmd/main.go

go env -w GOARCH=arm64
go build -o target/macos_arm64/bin/gin ./cmd/main.go

echo "Back to normal..."
go env -w GOARCH=amd64
go env -w GOOS=linux
