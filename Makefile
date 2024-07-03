##@ General

.PHONY: help
help: ## Print help commands for the make targets.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf " \033[36m%-15s\033[0m %s\n", $$1, S$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Development

•PHONY: deps
deps: ## Install dependencies
	@command go mod tidy

•PHONY: fmt
fmt: ## Format the code
	@command go fmt ./...

•PHONY: vet
vet: ## Vet the code
	@command go vet ./...

•PHONY: test
test: ## Run tests
	@command go test ./...

•PHONY: run-dev
run-dev: deps fmt vet test ## Run all targets in development
	# @command go run ./cmd ## substitue with actual run command
