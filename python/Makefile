PROJECT_NAME     := Khulnasoft Python SDK
LANGHOST_PKG     := github.com/khulnasoft/sdk/python/cmd/khulnasoft-language-python/v3
VERSION          := $(if ${KHULNASOFT_VERSION},${KHULNASOFT_VERSION},$(shell ../../scripts/khulnasoft-version.sh))
PYPI_VERSION 	   := $(if ${PYPI_VERSION},${PYPI_VERSION},$(shell ../../scripts/khulnasoft-version.sh python))

ifeq ($(DEBUG),"true")
$(info    VERSION       = $(VERSION))
$(info    PYPI_VERSION  = $(PYPI_VERSION))
endif

PYENV := ./env
PYENVSRC := $(PYENV)/src

PROJECT_PKGS    := $(shell go list -C cmd/khulnasoft-language-python ./...)

include ../../build/common.mk

# Motivation: running `make TEST_ALL_DEPS= test_all` permits running
# `test_all` without the dependencies.
TEST_ALL_DEPS ?= build

ensure:: $(PYTHON).ensure .ensure.phony
.ensure.phony:
	$(PYTHON) -m venv venv
	. venv/*/activate && python -m pip install --upgrade pip setuptools wheel
	. venv/*/activate && python -m pip install -r requirements.txt
	mkdir -p $(PYENVSRC)
	@touch .ensure.phony

build_package:: ensure
	rm -rf $(PYENVSRC) && cp -R ./lib/. $(PYENVSRC)/
	sed -i.bak 's/^VERSION = .*/VERSION = "$(PYPI_VERSION)"/g' $(PYENVSRC)/setup.py && rm $(PYENVSRC)/setup.py.bak
	sed -i.bak 's/^_VERSION = .*/_VERSION = "$(VERSION)"/g' $(PYENVSRC)/khulnasoft/_version.py && rm $(PYENVSRC)/khulnasoft/_version.py.bak
	cp ../../README.md $(PYENVSRC)
	. venv/*/activate && cd $(PYENVSRC) && \
		python setup.py build bdist_wheel --universal
	. venv/*/activate && python -m pip install -e $(PYENVSRC)

build_plugin::
	go install -C cmd/khulnasoft-language-python \
		-ldflags "-X github.com/khulnasoft/sdk/go/common/version.Version=${VERSION}" ${LANGHOST_PKG}

build:: build_package build_plugin

lint:: ensure
	. venv/*/activate && \
		python -m ruff format --check
	. venv/*/activate && \
		MYPYPATH=./stubs python -m mypy ./lib/khulnasoft --config-file=mypy.ini
	. venv/*/activate && \
		python -m ruff check
	. venv/*/activate && \
		python -m pyright

lint_fix:: ensure
	make format
	make lint

format:: ensure
	. venv/*/activate && \
		python -m ruff format

install_package:: build_package
	cp ./cmd/khulnasoft-language-python-exec "$(KHULNASOFT_BIN)"
	cp ./dist/khulnasoft-resource-khulnasoft-python "$(KHULNASOFT_BIN)"
	cp ./dist/khulnasoft-analyzer-policy-python "$(KHULNASOFT_BIN)"

install_plugin:: build_plugin
	GOBIN=$(KHULNASOFT_BIN) go install -C cmd/khulnasoft-language-python \
		  -ldflags "-X github.com/khulnasoft/sdk/go/common/version.Version=${VERSION}" ${LANGHOST_PKG}

install:: install_package install_plugin

test_go:: $(TEST_ALL_DEPS)
	@cd cmd/khulnasoft-language-python && \
		$(GO_TEST) ${PROJECT_PKGS}

test_fast:: $(TEST_ALL_DEPS)
	. venv/*/activate && ./scripts/test_fast.sh

test_auto:: $(TEST_ALL_DEPS)
	. venv/*/activate && ./scripts/test_auto.sh

test_all:: test_fast test_auto test_go

dist::
	go install -C cmd/khulnasoft-language-python \
		-ldflags "-X github.com/khulnasoft/sdk/go/common/version.Version=${VERSION}" ${LANGHOST_PKG}
	cp ./cmd/khulnasoft-language-python-exec "$$(go env GOPATH)"/bin/
	cp ./dist/khulnasoft-resource-khulnasoft-python "$$(go env GOPATH)"/bin/
	cp ./dist/khulnasoft-analyzer-policy-python "$$(go env GOPATH)"/bin/

brew:: BREW_VERSION := $(shell ../../scripts/get-version HEAD)
brew::
	go install -C cmd/khulnasoft-language-python \
		-ldflags "-X github.com/khulnasoft/sdk/go/common/version.Version=${BREW_VERSION}" ${LANGHOST_PKG}
	cp ./cmd/khulnasoft-language-python-exec "$$(go env GOPATH)"/bin/
	cp ./dist/khulnasoft-resource-khulnasoft-python "$$(go env GOPATH)"/bin/
	cp ./dist/khulnasoft-analyzer-policy-python "$$(go env GOPATH)"/bin/

publish::
	for file in ../../artifacts/sdk-python-*.whl; do \
		basename=$$(basename $${file}); \
		mv -vT "$${file}" "../../artifacts/$${basename##sdk-python-}"; \
	done
	twine upload \
	-u "${PYPI_USERNAME}" -p "${PYPI_PASSWORD}" \
		../../artifacts/*.whl \
		--skip-existing \
		--verbose
