.PHONY: all test examples

all: test

test: unitTest

unitTest:
	go test ${TEST_OPTS} ./...
	
examples:
	go run examples/*.go
