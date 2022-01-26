#!/bin/bash

# jenkins 的持续集成
# 因为使用来热重载，将不再构建 windows 版本

echo "Local dir $PWD"

echo "Setting goproxy..."

/usr/local/go/bin/go env -w GOPROXY=https://goproxy.cn,direct
/usr/local/go/bin/go env -w GO111MODULE=on
/usr/local/go/bin/go mod tidy

echo "Testing ..."
/usr/local/go/bin/go env -w GOARCH=amd64
/usr/local/go/bin/go env -w GOOS=linux
/usr/local/go/bin/go test -v test/*.go

echo "Building package ..."

echo "Build linux..."
/usr/local/go/bin/go env -w GOARCH=amd64
/usr/local/go/bin/go env -w GOOS=linux
/usr/local/go/bin/go build -o target/linux/bin/gin ./cmd/main.go

echo "Build macos..."
/usr/local/go/bin/go env -w GOOS=darwin
/usr/local/go/bin/go build -o target/macos_amd64/bin/gin ./cmd/main.go

/usr/local/go/bin/go env -w GOARCH=arm64
/usr/local/go/bin/go build -o target/macos_arm64/bin/gin ./cmd/main.go

echo "Back to normal..."
/usr/local/go/bin/go env -w GOARCH=amd64
/usr/local/go/bin/go env -w GOOS=linux
