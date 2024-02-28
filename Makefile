# VARIABLES
PACKAGE="github.com/hasura/v3-cli-plugin-upgrade-from-v2"

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


####### Sections adapted from https://github.com/hasura/connectors-cloud-integration/blob/main/cli/Makefile

VERSION   ?= $(shell ./scripts/get-version.sh)
OS        ?= linux darwin windows

HAS_GOX   := $(shell command -v gox;)

BUILDDIR  := dist
ASSETS    := $(BUILDDIR)/hasura-upgrade-from-v2-darwin-amd64.tar.gz $(BUILDDIR)/hasura-upgrade-from-v2-darwin-arm64.tar.gz $(BUILDDIR)/hasura-upgrade-from-v2-linux-amd64.tar.gz $(BUILDDIR)/hasura-upgrade-from-v2-linux-arm64.tar.gz $(BUILDDIR)/hasura-upgrade-from-v2-windows-amd64.zip
CHECKSUMS := $(patsubst %,%.sha256,$(ASSETS))
COMPRESS  := gzip --best -k -c

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
	-output="$(BUILDDIR)/hasura-upgrade-from-v2-{{.OS}}-{{.Arch}}" \
	.

.PHONY: manifest
manifest: $(CHECKSUMS)
	./scripts/generate-manifest.sh && \
	rm $(BUILDDIR)/manifest-tmp.yaml

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

.PHONY: setup-gcloud
setup-gcloud:
	echo ${UPGRADE_FROM_V2_GCLOUD_SERVICE_KEY} | base64 --decode > ${HOME}/gcloud-service-key.json
	gcloud auth activate-service-account --key-file=${HOME}/gcloud-service-key.json
	gcloud --quiet config set connectorject ${UPGRADE_FROM_V2_GCLOUD_PROJECT_ID}

push-artifacts:
	# cp -r /build/_connector_cli_output/binaries ${BUILDDIR}
	# cp /build/_connector_cli_output/manifest-dev.yaml ${BUILDDIR}/manifest.yaml
	gsutil -m cp $(ASSETS) $(CHECKSUMS) gs://hasura-pro-cdn/hasura-upgrade-from-v2/$(VERSION)/

plugin-index-pr:
	$(PWD)/scripts/make-plugin-index-pr.sh
