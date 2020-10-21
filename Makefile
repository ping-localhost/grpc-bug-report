.PHONY: all deps run

all: run

deps: ## Installing all dependencies
	@cd server && go mod vendor

run: deps ## Starts the server
	 @cd server && go run server.go

help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
