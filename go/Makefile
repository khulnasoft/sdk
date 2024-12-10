PROJECT_NAME     := Khulnasoft Go SDK
LANGHOST_PKG     := github.com/khulnasoft/sdk/go/khulnasoft-language-go/v3
VERSION          := $(if ${KHULNASOFT_VERSION},${KHULNASOFT_VERSION},$(shell ../../scripts/khulnasoft-version.sh go))
TEST_FAST_PKGS   := $(shell go list ./khulnasoft/... ./common/... | grep -v /vendor/ | grep -v templates)
TEST_AUTO_PKGS   := $(shell go list ./auto/... | grep -v /vendor/ | grep -v templates)

ifeq ($(DEBUG),"true")
$(info    VERSION         = $(VERSION))
endif

include ../../build/common.mk

# Motivation: running `make TEST_ALL_DEPS= test_all` permits running
# `test_all` without the dependencies.
TEST_ALL_DEPS ?= install

gen::
	go generate ./khulnasoft/...

ensure: go.ensure

build:: gen
	go install -C khulnasoft-language-go \
		-ldflags "-X github.com/khulnasoft/sdk/go/common/version.Version=${VERSION}" ${LANGHOST_PKG}

install_plugin::
	GOBIN=$(KHULNASOFT_BIN) go install -C khulnasoft-language-go \
		-ldflags "-X github.com/khulnasoft/sdk/go/common/version.Version=${VERSION}" ${LANGHOST_PKG}

install:: install_plugin

test_all:: test_fast test_auto

test_fast:: $(TEST_ALL_DEPS)
	@$(GO_TEST_FAST) $(TEST_FAST_PKGS)

	@cd khulnasoft-language-go && \
		$(GO_TEST_FAST) $(shell cd khulnasoft-language-go && go list ./... | grep -v /vendor/ | grep -v templates)

test_auto:: $(TEST_ALL_DEPS)
	@$(GO_TEST) $(TEST_AUTO_PKGS)

dist::
	go install -C khulnasoft-language-go \
		-ldflags "-X github.com/khulnasoft/sdk/go/common/version.Version=${VERSION}" ${LANGHOST_PKG}

brew:: BREW_VERSION := $(shell ../../scripts/get-version HEAD)
brew::
	go install -C khulnasoft-language-go \
		-ldflags "-X github.com/khulnasoft/sdk/go/common/version.Version=${BREW_VERSION}" ${LANGHOST_PKG}

lint:: golangci-lint.ensure
	cd .. && golangci-lint run -c ../.golangci.yml --timeout 5m --path-prefix ..

	cd khulnasoft-language-go && \
		golangci-lint run -c ../../../.golangci.yml --timeout 5m --path-prefix khulnasoft-language-go

publish:
	git tag sdk/v${VERSION}
	git push origin sdk/v${VERSION}
