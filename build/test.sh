#!/bin/sh

set -o errexit
set -o nounset
set -o pipefail

echo "Running tests:"
go test -v .
echo
