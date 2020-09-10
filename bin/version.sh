#!/usr/bin/env bash
set -euo pipefail

grep -E "## \[[0-9]*.[0-9]*.[0-9]*\]" CHANGELOG.md | head -1 | cut -d "[" -f2 | cut -d "]" -f1
