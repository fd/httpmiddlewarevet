#!/usr/bin/env bash

set -e

GIMME_OUTPUT="$(gimme 1.9.x | tee -a ${TRAVIS_HOME}/.bashrc)" && eval "$GIMME_OUTPUT"

echo "$GIMME_OUTPUT"

GO_VERSION=$(echo "$GIMME_OUTPUT" | awk '{ print $3 }')

echo "$GO_VERSION"

GO_VERSION=${GO_VERSION#"go"}

echo "$GO_VERSION"

curl -X GET "https://httpmiddlewarevet.herokuapp.com/webhook/travis?commit=${TRAVIS_COMMIT}&version=${GO_VERSION}&branch=${TRAVIS_BRANCH}"
