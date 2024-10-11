# ==================================================================================== #
# HELPERS
# ==================================================================================== #

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

.PHONY: confirm
confirm:
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]

# ==================================================================================== #
# QUALITY CONTROL
# ==================================================================================== #

## test: run all tests
.PHONY: test
test:
	go test -v ./...

# ==================================================================================== #
# BUILD
# ==================================================================================== #

## build/charming
## build/charming:
.PHONY: build/charming
build/charming:
	make test

	@echo 'Building charming...'
	go build -o ./bin/charming ./cmd/charming/
