exe = cmd/operator/*
cmd = operator
TRAVIS_TAG ?= "0.0.0"

.PHONY: deps lint test integration integration-windows git-hooks init

init: git-hooks

git-hooks:
	$(info INFO: Starting build $@)
	ln -sf ../../.githooks/pre-commit .git/hooks/pre-commit

deps:
	$(info INFO: Starting build $@)
	go mod vendor

build:
	$(info INFO: Starting build $@)
	go build $(exe)

lint:
	$(info INFO: Starting build $@)
	golint pkg/ cmd/

test:
	$(info INFO: Starting build $@)
	go test ./...

test-coverage:
	$(info INFO: Starting build $@)
	go test -coverprofile c.out ./...

release-amd64:
	$(info INFO: Starting build $@)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-X main.version=$(TRAVIS_TAG) -s -w" -o release/$(cmd)-linux-amd64 $(exe)

release-arm:
	$(info INFO: Starting build $@)
	CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -ldflags "-X main.version=$(TRAVIS_TAG) -s -w" -o release/$(cmd)-linux-arm $(exe)

release-386:
	$(info INFO: Starting build $@)
	CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -ldflags "-X main.version=$(TRAVIS_TAG) -s -w" -o release/$(cmd)-linux-386 $(exe)

release: release-amd64 release-arm release-386 release-darwin-amd64 release-darwin-386 release-windows-amd64 release-windows-386