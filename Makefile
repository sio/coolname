GO?=go
GIT?=git

.PHONY: todo
todo:
	@$(GIT) grep TODO -- ':(exclude)Makefile'

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

.PHONY: ci
ci: versions codegen lint test bench

.PHONY: reset-upstream-ref
reset-upstream-ref:
	echo > data/codegen/upstream.ref

.PHONY: versions
versions:
	@$(GO) version
	@$(GIT) --version
