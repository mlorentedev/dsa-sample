# Variables
PROJECT_NAME = algorithms
SRC_DIR = core
TEST_DIR = tests
BIN_DIR = bin

.PHONY: help
help:  ## Show this help | make help
	@grep -E '^\S+:.*?## .*$$' $(firstword $(MAKEFILE_LIST)) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "%-30s %s\n", $$1, $$2}'

.PHONY: todo
todo:  ## Show the TODOs in the code | make todo
	@{ grep -n -w TODO Makefile | uniq || echo "No pending tasks"; } | sed '/grep/d'

.PHONY: init
init: ## Initialize the project | make init
	@echo "Initializing the project $(PROJECT_NAME)"
	@go mod init $(PROJECT_NAME)

.PHONY: install $(DEPENDENCY)
install: ## Install some dependency | make install DEPENDENCY=github.com/dependency
	@echo "Installing the dependency $(DEPENDENCY)"
	@go get -u $(DEPENDENCY)

.PHONY: get-dependencies
get-dependencies: ## Get the dependencies of the project | make get-dependencies
	@echo "Getting the dependencies of the project $(PROJECT_NAME)"
	@go get -v ./...

.PHONY: update-dependencies
update-dependencies: ## Update the dependencies of the project | make update-dependencies
	@echo "Updating the dependencies of the project $(PROJECT_NAME)"
	@go get -u ./...

.PHONY: clean
clean: ## Clean the project | make clean
	@echo "Cleaning the project $(PROJECT_NAME)"
	@rm -rf $(BIN_DIR) $(TEST_DIR)/*.out $(TEST_DIR)/*.html
	@go mod tidy

.PHONY: build
build: ## Build the project | make build
	@echo "Building the project $(PROJECT_NAME) in $(BIN_DIR)"
	@go build -o $(BIN_DIR)/$(PROJECT_NAME) .

.PHONY: run
run: ## Run the project | make run
	@echo "Running the project $(PROJECT_NAME)"
	@$(BIN_DIR)/$(PROJECT_NAME)

.PHONY: reformat
reformat: ## Reformat the code | make reformat
	@echo "Reformatting the code"
	@go fmt ./...

.PHONY: check-typing
check-typing: ## Check the typing of the code | make check-typing
	@echo "Checking the typing of the code"
	@go vet ./...

.PHONY: pre-commit
pre-commit: reformat check-typing
	@echo "Running the pre-commit checks"
	$(MAKE) reformat
	$(MAKE) check-typing
	
.PHONY: test-unit
test-unit: ## Run the unit tests | make test-unit
	@echo "Running the unit tests"
	@go test -v -cover -coverpkg=./$(SRC_DIR) -coverprofile=$(TEST_DIR)/coverage.out ./$(TEST_DIR)/unit
	@go tool cover -html=$(TEST_DIR)/coverage.out -o $(TEST_DIR)/coverage.html

.PHONY: test-load
test-load: ## Run the load tests | make test-load
	@go test -bench=. ./$(TEST_DIR)/load

