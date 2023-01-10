GO?=go

.PHONY: test
test:
	$(GO) test -race -timeout 15s ./...

.PHONY: lint
lint: fmt
	$(GO) vet ./...

.PHONY: fmt
fmt:
	$(GO) fmt ./...

.PHONY: codegen
codegen:
	$(GO) generate ./...
