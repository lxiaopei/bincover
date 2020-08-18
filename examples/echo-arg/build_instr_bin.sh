#!/usr/bin/env bash
go test . -tags testbincover -coverpkg=./... -c -o instr_bin -ldflags="-X github.com/lxiaopei/bincover/examples/echo-arg.isTest=true"
