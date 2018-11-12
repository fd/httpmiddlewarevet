#!/usr/bin/env bash

set -e

GO_VERSION=$(gimme --resolve "$TRAVIS_GO_VERSION")

echo "$GO_VERSION"

curl -X GET "https://httpmiddlewarevet.herokuapp.com/webhook/travis?commit=${TRAVIS_COMMIT}&version=${GO_VERSION}&branch=${TRAVIS_BRANCH}"
