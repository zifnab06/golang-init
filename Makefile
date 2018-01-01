.PHONY: build clean

build: clean init

deps:
	go get github.com/gliderlabs/ssh
	go get golang.org/x/crypto/ssh

clean:
	-rm init

init:
	CGO_ENABLED=0
	go build -a -installsuffix cgo -ldflags '-s' init.go config.go http.go ssh.go network.go
