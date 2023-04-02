#!/usr/bin/env bash

go version || exit 9

# Linux
echo "Linux"                                                       && \
GOOS=linux GOARCH=amd64 go build -o aegis-cli-linux-x86_64 main.go && \
GOOS=linux GOARCH=arm64 go build -o aegis-cli-linux-arm64  main.go && \

# macOS
echo "macOS"                                                        && \
GOOS=darwin GOARCH=amd64 go build -o aegis-cli-macos-x86_64 main.go && \
GOOS=darwin GOARCH=arm64 go build -o aegis-cli-macos-arm64  main.go && \

echo "✅ DONE"
