.PHONY: build clean

build: 
	export GO111MODULE=on
	#GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" -o ./bin/player main.go
	go build -o bin/player main.go

clean:
	rm -Rf bin
