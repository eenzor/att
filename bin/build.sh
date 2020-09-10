#!/usr/bin/env bash
set -euo pipefail

COMMIT=$(git rev-parse --short HEAD)
VERSION=$(./bin/version.sh)

echo "building version: $VERSION from commit: $COMMIT"

CGO_ENABLED=0 go build -ldflags "-X main.version=$VERSION -X main.commit=$COMMIT" -a -o att main.go
