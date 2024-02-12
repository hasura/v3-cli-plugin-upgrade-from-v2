# VARIABLES
PACKAGE="upgrade-from-v2"

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
