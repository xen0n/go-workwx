#!/bin/bash

set -e

TOP="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "${TOP}"


die() {
    echo fatal: lint failed: $@ >&2
    exit 1
}


# gofmt conformance check
check_gofmt() {
    # record the list of non-conforming source files
    local output="$(gofmt -s -l $@)"
    if [[ "x${output}" != "x" ]]; then
        # report the offending files then die
        die "the following files violate gofmt -s style: ${output}"
    fi
}

# exclude vendored files
# TODO: write this more elegantly
check_gofmt *.go cmd


# staticcheck for known problems
staticcheck . || die "staticcheck reported error(s)"
