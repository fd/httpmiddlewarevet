#!/usr/bin/env bash

set -e

GIMME_OUTPUT="go version go1.11beta1 linux/amd64"
GO_VERSION=$(echo "$GIMME_OUTPUT" | awk '{ print $3 }')
GO_VERSION=${GO_VERSION#"go"}

echo "$GO_VERSION"

curl -X GET "https://httpmiddlewarevet.herokuapp.com/webhook/travis?commit=${TRAVIS_COMMIT}&version=${GO_VERSION}&branch=${TRAVIS_BRANCH}"
