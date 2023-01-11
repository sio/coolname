GO?=go

.PHONY: test
test:
	$(GO) test -timeout 15s ./...

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
