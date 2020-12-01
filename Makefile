NAME?=pokemon
ARCH=amd64
BIN = bin/pokemon
BIN_LINUX = $(BIN)-linux-$(ARCH)
BIN_DARWIN = $(BIN)-darwin-$(ARCH)
GO111MODULE=on
GIT_BRANCH?=$(shell git rev-parse --abbrev-ref HEAD)
IMG_NAMESPACE?=quay.io/hriships
IMG_TAG?=$(GIT_BRANCH)
REGISTRY?=$(IMG_NAMESPACE)/$(NAME)

.PHONEY: unit-test all clean

all: $(BIN_LINUX) $(BIN_DARWIN)

$(BIN_DARWIN):
	GOARCH=$(ARCH) GOOS=darwin CGO_ENABLED=0 go build -o $(BIN_DARWIN) cmd/main.go

$(BIN_LINUX):
	GOARCH=$(ARCH) GOOS=linux CGO_ENABLED=0 go build -o $(BIN_LINUX) cmd/main.go

unit-test:
	go test ./pkg/...

integration-test:
	go test ./tests -v

docker: Dockerfile
	docker image build -t "$(REGISTRY):$(IMG_TAG)" .

clean:
	rm -rf bin/