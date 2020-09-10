#!/usr/bin/env bash
set -euo pipefail
VERSION=$(./bin/version.sh)

docker build .  --tag att:${VERSION} --tag att:latest