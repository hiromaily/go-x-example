# Note: tabs by space can't not used for Makefile!

CURRENTDIR=`pwd`

###############################################################################
# Managing Dependencies
###############################################################################
.PHONY: update
update:
	GO111MODULE=off go get -u github.com/golangci/golangci-lint/cmd/golangci-lint
	go get -u -d -v ./...


###############################################################################
# Golang formatter and detection
###############################################################################
.PHONY: lint
lint:
	golangci-lint run --fix

.PHONY: imports
imports:
	./scripts/imports.sh


###############################################################################
# Test run
###############################################################################

test-eg:
	go test -v -run=ErrorGroup ./sync/...

test-eg2:
	go test -v -run=Parallel ./sync/...

test-eg3:
	go test -v -run=ParallelWithTimeout ./sync/...
