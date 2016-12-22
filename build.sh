#!/bin/bash

set -x

if [ -z $GOPATH ]; then
    export GOPATH=`pwd`
fi

go get github.com/hashicorp/go-plugin \
&& go fmt lahenry.com/... \
&& go install lahenry.com/...

