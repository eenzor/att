#!/usr/bin/env bash
COMMIT=$(git rev-parse --short HEAD)
VERSION="dev"
if [[ $# -eq 1 ]]; then
    VERSION=$1
fi

echo "building version: $VERSION from commit: $COMMIT"

go build -ldflags "-X main.ver=$VERSION -X main.commit=$COMMIT" -o att main.go
