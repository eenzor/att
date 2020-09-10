#!/usr/bin/env bash
set -euo pipefail

COMMIT=$(git rev-parse --short HEAD)
VERSION=$(grep -E "## \[[0-9]*.[0-9]*.[0-9]*\]" CHANGELOG.md | head -1 | cut -d "[" -f2 | cut -d "]" -f1 )

echo "building version: $VERSION from commit: $COMMIT"

go build -ldflags "-X main.ver=$VERSION -X main.commit=$COMMIT" -o att main.go
