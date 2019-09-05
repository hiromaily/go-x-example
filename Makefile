.PHONY: imports
imports:
	./scripts/rm-blank-line.sh

test-eg:
	go test -v -run=ErrorGroup ./sync/...

test-eg2:
	go test -v -run=Parallel ./sync/...

test-eg3:
	go test -v -run=ParallelWithTimeout ./sync/...