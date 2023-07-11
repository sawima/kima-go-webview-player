.PHONY: build clean

build: 
	export GO111MODULE=on
	#GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" -o ./bin/player main.go
	# go build -o bin/player main.go
	CGO_ENABLED=1 GOOS=linux GOARCH=arm64 CC_FOR_TARGET=gcc-aarch64-linux-gnu GOOS=linux GOARCH=arm64 go build -o bin/player main.go

clean:
	rm -Rf bin
