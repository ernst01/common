# Go parameters
GOCMD=go

MAKEFLAGS += --silent

## test: Runs your tests if any
test:
	@echo " > Running tests..."
	$(GOCMD) test ./... -cover -v -coverprofile=coverage.out
	@echo " > Done."

## cover: Checks your code coverage
cover:
	@echo " > Checking coverage..."
	$(GOCMD) tool cover -html=coverage.out

.PHONY: help
all: help
help: Makefile
	@echo
	@echo "Choose a command to run in "$(PROJECTNAME)":"
	@echo
	sed -n 's/^##//p' $< | column -t -s ':' | sed -e 's/^/ /'
	@echo