GITCOMMIT := $(shell git rev-parse HEAD)
GITDATE := $(shell git show -s --format='%ct')

LDFLAGSSTRING +=-X main.GitCommit=$(GITCOMMIT)
LDFLAGSSTRING +=-X main.GitDate=$(GITDATE)
LDFLAGS := -ldflags "$(LDFLAGSSTRING)"

bridge:
	env GO111MODULE=on go build -v $(LDFLAGS) ./cmd/bridge

test:
	go test -v ./...

abigen:
	bash ./scripts/abigen.sh

lint:
	golangci-lint run ./...

.PHONY: \
	abigen \
	bridge \
	clean \
	test \
	lint