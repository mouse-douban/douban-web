#!/bin/bash

# jenkins 的持续集成
# 因为使用来热重载，将不再构建 windows 版本

echo "Local dir $PWD"

echo "Setting goproxy..."

$GOROOT/bin/go env -w GOPROXY=https://goproxy.cn,direct
$GOROOT/bin/go env -w GO111MODULE=on
$GOROOT/bin/go mod tidy

echo "Testing ..."

$GOROOT/bin/go test -v test/*.go

echo "Building package ..."

echo "Build linux..."
$GOROOT/bin/go env -w GOARCH=amd64
$GOROOT/bin/go env -w GOOS=linux
$GOROOT/bin/go build -o target/linux/bin/gin ./cmd/main.go

echo "Build macos..."
$GOROOT/bin/go env -w GOOS=darwin
$GOROOT/bin/go build -o target/macos_amd64/bin/gin ./cmd/main.go

$GOROOT/bin/go env -w GOARCH=arm64
$GOROOT/bin/go build -o target/macos_arm64/bin/gin ./cmd/main.go

echo "Back to normal..."
$GOROOT/bin/go env -w GOARCH=amd64
$GOROOT/bin/go env -w GOOS=linux
