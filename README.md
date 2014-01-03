Registry-server
===============

This is a centralized authentication and session management server intended to fulfill the features of a SSO server. Also check [registry-client](https://github.com/eduvo/registry-client) for rails integration.

This code is very young and still work in progress, don't use it !

## Install

1. `go get github.com/mattn/gom`
1. `gom install`

## Launch

    find . -d 1 -name '*.go' -not -name '*_test.go' | xargs gom run

or add an alias in your .bashrc or .zshrc, it would work on any gom-based project

    alias grun="find . -d 1 -name '*.go' -not -name '*_test.go' | xargs gom run"

## Create ssl certs

For dev purposes

    go run $GOROOT/src/pkg/crypto/tls/generate_cert.go --host="localhost"

For production better have real certificates.

## Copyright

Copyright 2013 Faria Systems Ltd.
Licensed under MIT
