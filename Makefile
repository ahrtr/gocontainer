.PHONY: all test

all: test

test: unitTest

unitTest:
	go test ${TEST_OPTS} github.com/ahrtr/gocontainer/stack
	go test ${TEST_OPTS} github.com/ahrtr/gocontainer/queue
	go test ${TEST_OPTS} github.com/ahrtr/gocontainer/set
	go test ${TEST_OPTS} github.com/ahrtr/gocontainer/list
	go test ${TEST_OPTS} github.com/ahrtr/gocontainer/priorityqueue
	go test ${TEST_OPTS} github.com/ahrtr/gocontainer/sort
	
