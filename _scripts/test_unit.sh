#!/bin/sh
go test -v -cover ./... | grep -E "PASS|FAIL|coverage"