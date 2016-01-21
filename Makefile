.PHONY: test

all: check test

check: goimports govet golint

goimports:
	@echo checking go imports...
	@! goimports -d . 2>&1 | egrep -v '^$$'

govet:
	@echo checking go vet...
	@go tool vet -structtags=false -methods=false .

golint:
	@echo checking go lint ...
	@go get -v github.com/golang/lint/golint
	@for file in $$(find . -name '*.go' | grep -v 'vendor\|govc\|vim25\|test'); do \
		golint $${file}; \
		if [ -n "$$(golint $${file})" ]; then \
			exit 1; \
		fi; \
	done


test:
	go get
	go test -v $(TEST_OPTS) ./...

install:
	go install github.com/vmware/govmomi/govc
