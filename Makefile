GO?=go
GIT?=git

export GOTEST_ARGS

.PHONY: todo
todo:
	@$(GIT) grep $(addsuffix DO,TO) # avoid self-reference in grep

.PHONY: test
test:
	$(GO) test -timeout=15s $(GOTEST_ARGS) ./...

.PHONY: test-verbose
test-verbose: GOTEST_ARGS+=-v
test-verbose:
	@$(MAKE) test

.PHONY: test-multi
test-multi: GOTEST_ARGS+=-count=1000
test-multi: GOTEST_ARGS+=-short
test-multi:
	@$(MAKE) test

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
ci: versions codegen lint test bench test-multi

.PHONY: reset-upstream-ref
reset-upstream-ref:
	echo > data/codegen/upstream.ref

.PHONY: versions
versions:
	@$(GO) version
	@$(GIT) --version
