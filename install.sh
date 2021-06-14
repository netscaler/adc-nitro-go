#!/bin/sh
# This script will copy the lib to $GOPATH/src/github.com/citrix/adc-nitro-go
mkdir -p $GOPATH/src/github.com/citrix/adc-nitro-go
cp -r * $GOPATH/src/github.com/citrix/adc-nitro-go
