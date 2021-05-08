GO                      ?= go
pkgs                    = $(shell $(GO) list ./... | grep -v 'github.com/vmware/govmomi/vim25/xml')

all: check test

check: goimports govet

goimports:
	@echo checking go imports...
	@command -v goimports >/dev/null 2>&1 || $(GO) get golang.org/x/tools/cmd/goimports
	@! goimports -d . 2>&1 | egrep -v '^$$'
	@! TERM=xterm git grep encoding/xml -- '*.go' ':!vim25/xml/*.go'

govet:
	@echo checking go vet...
	@$(GO) vet -structtag=false -methods=false $(pkgs)

install:
	$(MAKE) -C govc install
	$(MAKE) -C vcsim install

go-test:
	GORACE=history_size=5 $(GO) test -mod=vendor -timeout 5m -count 1 -race -v $(TEST_OPTS) ./...

govc-test: install
	./govc/test/images/update.sh
	(cd govc/test && ./vendor/github.com/sstephenson/bats/libexec/bats -t .)

govc-test-sso: install
	./govc/test/images/update.sh
	(cd govc/test && SSO_BATS=1 ./vendor/github.com/sstephenson/bats/libexec/bats -t sso.bats)

govc-test-sso-assert-cert:
	SSO_BATS_ASSERT_CERT=1 $(MAKE) govc-test-sso

.PHONY: test
test: go-test govc-test

doc: install
	./govc/usage.sh > ./govc/USAGE.md