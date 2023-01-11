GO?=go

.PHONY: todo
todo:
	@git grep TODO -- ':(exclude)Makefile'

.PHONY: test
test:
	$(GO) test -timeout 15s ./...

.PHONY: bench
bench:
	$(GO) test -bench=. -count=3 -benchmem -benchtime=2s -run='^#' ./...

.PHONY: lint
lint: fmt
	$(GO) vet ./...
	git diff --exit-code --name-only

.PHONY: fmt
fmt:
	$(GO) fmt ./...

.PHONY: codegen
codegen:
	$(GO) generate ./...
