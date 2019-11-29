#!/bin/sh

echo "Checking gofmt: "
ERRS=$(gofmt -l .)
if [ -n "${ERRS}" ]; then
    echo "FAIL - the following files need to be formatted:"
    for e in ${ERRS}; do
        echo "    $e"
    done
    echo
    exit 1
fi
echo "PASS"
echo
