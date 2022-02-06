#!/bin/bash

echo "Local dir $PWD"

echo "Set local config loading..."
sed -i 's/EnableCOS = true/EnableCOS = false/g' cmd/main.go

echo "Setting goproxy..."

/usr/local/go/bin/go env -w GOPROXY=https://goproxy.cn,direct
/usr/local/go/bin/go env -w GO111MODULE=on
/usr/local/go/bin/go mod tidy

echo "Testing ..."
/usr/local/go/bin/go env -w GOARCH=amd64
/usr/local/go/bin/go env -w GOOS=linux
/usr/local/go/bin/go test -v test/*/*.go

echo "Building package ..."

echo "Build linux..."
/usr/local/go/bin/go env -w GOARCH=amd64
/usr/local/go/bin/go env -w GOOS=linux
/usr/local/go/bin/go build -o target/linux_gin ./cmd/main.go

echo "Build darwin..."
/usr/local/go/bin/go env -w GOOS=darwin
/usr/local/go/bin/go build -o target/darwin_amd64_gin ./cmd/main.go

/usr/local/go/bin/go env -w GOARCH=arm64
/usr/local/go/bin/go build -o target/darwin_arm64_gin ./cmd/main.go

echo "Back to normal..."
/usr/local/go/bin/go env -w GOARCH=amd64
/usr/local/go/bin/go env -w GOOS=linux
