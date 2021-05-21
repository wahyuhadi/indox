test:
	go test -v ./...

build:
	go build

install:
	go install

compile:
	echo "Compiling for every OS and Platform"
	GOOS=linux GOARCH=arm go build -o bin/main-linux-arm 
	GOOS=linux GOARCH=arm64 go build  -o bin/main-linux-arm64 
	GOOS=freebsd GOARCH=386 go build -o bin/main-freebsd-386

all: hello build