#!/bin/bash

set -e

TOP="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "${TOP}"


die() {
    echo fatal: lint failed: $@ >&2
    exit 1
}


# golangci-lint for known problems
golangci-lint run || die "staticcheck reported error(s)"
