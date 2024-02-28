# VARIABLES
PACKAGE="github.com/hasura/v3-cli-plugin-upgrade-from-v2"
VERSION ?= $(shell ./scripts/get-version.sh)
BUILDDIR := dist
OS ?= linux darwin windows
COMPRESS := gzip --best -k -c
LINUX_FILES := ls dist | grep linux
DARWIN_FILES := ls dist | grep darwin 
WINDOWS_FILES := ls dist | grep windows

.PHONY: default
default: usage

.PHONY: clean
clean: ## Trash binary files
	@echo "--> cleaning..."
	@go clean || (echo "Unable to clean project" && exit 1)
	@echo "Clean OK"

.PHONY: test
test: ## Run all tests
	@echo "--> testing..."
	@go test -v $(PACKAGE)/...

.PHONY: install
install: clean ## Compile sources and build binary
	@echo "--> installing..."
	@go install $(PACKAGE) || (echo "Compilation error" && exit 1)
	@echo "Install OK"

.PHONY: run
run: install ## Run your application
	@echo "--> running application..."
	@go run main.go

.PHONY: dev
dev:
	@echo "Starting dev loop..."
	find . -name '*.go' -o -name '*.md' | entr make run

.PHONY: usage
usage: ## List available targets
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

# build connector locally, for all given platform/arch
.PHONY: build
build: export CGO_ENABLED=0
build:
ifndef HAS_GOX
	# so that go.mod don't get editted
	go install github.com/mitchellh/gox@latest
endif
	gox -ldflags '-X github.com/hasura/v3-cli-plugin-upgrade-from-v2/cli/pkg/version.BuildVersion=$(VERSION) -s -w -extldflags "-static"' \
	-rebuild \
	-os="$(OS)" \
	-arch="amd64 arm64" \
	-output="$(BUILDDIR)/upgrade-from-v2-{{.OS}}-{{.Arch}}" \
	.

# compress
.PHONY: compress
compress:
	for i in $$(ls dist | grep linux); do make "dist/$$i.tar.gz.sha256"; done
	for i in $$(ls dist | grep darwin); do make "dist/$$i.tar.gz.sha256"; done
	for i in $$(ls dist | grep windows); do make "dist/$$(basename $$i .exe).zip.sha256"; done
	for i in $$(ls dist | grep sha256); do cat dist/$$i; done	

.PRECIOUS: %.zip
%.zip: %.exe
	cd $(BUILDDIR) && \
	zip $(patsubst $(BUILDDIR)/%, %, $@) $(patsubst $(BUILDDIR)/%, %, $<)

.PRECIOUS: %.gz
%.gz: %
	$(COMPRESS) "$<" > "$@"

%.tar: %
	tar cf "$@" -C $(BUILDDIR) $(patsubst $(BUILDDIR)/%,%,$^)

%.sha256: %
	shasum -a 256 $< > $@

