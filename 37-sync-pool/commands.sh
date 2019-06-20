#!/usr/bin/env bash
set -ex

t(){
	go test -v
}

b(){
	go test -v -benchmem -bench=.
}

$@