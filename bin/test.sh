#!/usr/bin/env bash

echo "### Running go vet ###"
go vet .

echo "### Running go test ###"
go test -v -race -timeout 30s

echo "### Running golangci-lint ###"
docker run --rm -v $(pwd):/app -w /app golangci/golangci-lint:v1.31.0 golangci-lint run -v
