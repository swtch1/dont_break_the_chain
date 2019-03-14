#!/usr/bin/env bash

# ensure we're in the correct directory
cd $(dirname $0)

go build -o dbtc -mod=vendor .
