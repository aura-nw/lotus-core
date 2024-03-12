GITCOMMIT := $(shell git rev-parse HEAD)
GITDATE := $(shell git show -s --format='%ct')

LDFLAGSSTRING +=-X main.GitCommit=$(GITCOMMIT)
LDFLAGSSTRING +=-X main.GitDate=$(GITDATE)
LDFLAGS := -ldflags "$(LDFLAGSSTRING)"

bridge:
	env GO111MODULE=on go build -v $(LDFLAGS) ./cmd/bridge

operator:
	env GO111MODULE=on go build -v $(LDFLAGS) ./cmd/operator

test:
	go test -v ./...

proto-gen:
	bash ./scripts/protoc.sh protos/envelope.proto

abigen:
	bash ./scripts/abigen.sh

lint:
	golangci-lint run ./...

.PHONY: \
	abigen \
	bridge \
	clean \
	test \
	lint \
	proto-gen
