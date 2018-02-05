#!/bin/bash

ROOT=$GOPATH/src/github.com/heroku/task
APPNAME=task
ENV=$ROOT/env/loc.env

function start {
	SENV=$1;SAPP=$2
	set -a;. ${SENV};set +a;${SAPP} &
}

trap 'killall ${APPNAME}' SIGINT
start ${ENV} ${APPNAME}

