#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

TRAVISBUILD=${TRAVIS:-}

if [ ! -z "${TRAVISBUILD}" ]; then
  echo "Updating golang dependencies"
  go get -u github.com/axw/gocov/gocov
  go get -u github.com/AlekSi/gocov-xml
  go get -u github.com/jstemmer/go-junit-report
  curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b $(go env GOPATH)/bin v1.15.0
fi

export GOROOT=$(go env GOROOT)
FILE="${FILE:-0}"

echo
echo "Running golangci"
if [ ${FILE} -eq 0 ]; then
	golangci-lint run --enable-all --disable=lll --disable=gocyclo --disable=dupl --disable=unparam --disable=gochecknoglobals --deadline=300s --tests=false ./...
else
	golangci-lint run --enable-all --disable=lll --disable=gocyclo --disable=dupl --disable=unparam --disable=gochecknoglobals --deadline=300s --tests=false ./... | tee /dev/tty > checkstyle-report.xml
fi

echo
export CGO_ENABLED=0
echo "Running tests:"
if [ ${FILE} -eq 0 ]; then
  go test -v ./...
else
  go test -v ./... | tee /dev/tty | go-junit-report > junit-report.xml
fi

echo
echo "Testing coverage"

go test -covermode=count -coverprofile=profile.cov ./...
go tool cover -func profile.cov

if [ ${FILE} -eq 0 ]; then
  rm -f profile.cov
else
  gocov convert profile.cov | gocov-xml > coverage.xml
fi
echo
