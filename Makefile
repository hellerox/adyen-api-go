SHELL:=/bin/bash

test:
	go test -covermode=count -coverprofile=./profile.coverprofile -parallel 5 -v ./...

verification:
	go vet ./...
	golangci-lint run --enable-all --tests=false ./...
	CGO_ENABLED=0 errcheck ./...
