ARCH=$(shell go env GOARCH)
OS=$(shell go env GOOS)

.ONESHELL:
build-bin:
	sed -i '' 's/EnableCOS = false/EnableCOS = true/g' cmd/main.go;
	sed -i '' 's/InitWithCOS("config.json")/InitWithCOS("config2.json")/g' cmd/main.go;

	go mod tidy;

	go env -w GOARCH=amd64;
	go env -w GOOS=linux;

	go build -o target/linux_amd64_douban-web ./cmd/main.go;

	go env -w GOARCH=arm64;
	go env -w GOOS=darwin;

	go build -o target/darwin_arm64_douban-web ./cmd/main.go;

	go env -w GOARCH=amd64;
	go env -w GOOS=darwin;

	go build -o target/darwin_amd64_douban-web ./cmd/main.go;

	go env -w GOARCH=$(ARCH);
	go env -w GOOS=$(OS)


.ONESHELL:
build-image:
	docker image rm -f douban-web;
	docker build -t douban-web .
