.PHONY: all test unitTest examples

all: test

test: unitTest examples

unitTest:
	go test ${TEST_OPTS} ./...
	
examples:
	go run examples/*.go
