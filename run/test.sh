#!/usr/bin/env bash

cd ..

set -e
echo "" > coverage.txt

for d in $(go list ./... | grep -v vendor); do
    vet_flags=""
    if [[ $d == *"/surebetSearch/chrome" ]]; then
        vet_flags="-lostcancel=false"
    fi

    go vet $vet_flags $d

    staticcheck $d

    go test -v -race -coverprofile=profile.out -covermode=atomic $d
    if [ -f profile.out ]; then
        cat profile.out >> coverage.txt
        rm profile.out
    fi
done