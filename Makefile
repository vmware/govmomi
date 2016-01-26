.PHONY: test

$(shell export GO15VENDOREXPERIMENT=1)

all: check gvt dependencies test

check: goimports govet

goimports:
	@echo checking go imports...
	@! goimports -d . 2>&1 | egrep -v '^$$'

govet:
	@echo checking go vet...
	@go tool vet -structtags=false -methods=false .

gvt:
	@echo getting gvt
	go get -u github.com/FiloSottile/gvt

dependencies:
	@echo restoring dependencies
	$(GOPATH)/bin/gvt restore

test:
	go test -v $(TEST_OPTS) ./...

install:
	go install github.com/vmware/govmomi/govc
