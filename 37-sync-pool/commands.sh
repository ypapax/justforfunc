#!/usr/bin/env bash
set -ex

t(){
	go test -v
}

$@