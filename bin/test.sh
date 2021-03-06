#!/usr/bin/env bash
set -euo pipefail

echo "### Linting pipeline ###"
yamllint .github/workflows/

echo "### Running go vet ###"
go vet .

echo "### Running golangci-lint ###"
docker run --rm -v $(pwd):/app -w /app golangci/golangci-lint:v1.31.0 golangci-lint run

echo "### Running go test ###"
go test -v -race -timeout 30s -cover

echo "### Running gosec security scanner ###"
docker run --rm -v $(pwd):/app -w /app securego/gosec ./...
