exe = main.go
cmd = monorepo-operator
VERSION ?= "0.0.0"
BUILD_NUMBER ?= dev
REF ?= dev

.PHONY: deps lint test integration integration-windows git-hooks init docker-image docker-push

init: git-hooks

git-hooks:
	$(info INFO: Starting build $@)
	ln -sf ../../.githooks/pre-commit .git/hooks/pre-commit

deps:
	$(info INFO: Starting build $@)
	go mod vendor

build:
	$(info INFO: Starting build $@)
	go build -o $(cmd) $(exe)

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
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-X main.version=$(VERSION) -X main.build=$(BUILD_NUMBER) -X main.ref=$(REF) -s -w" -o release/$(cmd)-linux-amd64 $(exe)

release-arm:
	$(info INFO: Starting build $@)
	CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -ldflags "-X main.version=$(VERSION) -X main.build=$(BUILD_NUMBER) -X main.ref=$(REF) -s -w" -o release/$(cmd)-linux-arm $(exe)

release-386:
	$(info INFO: Starting build $@)
	CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -ldflags "-X main.version=$(VERSION) -X main.build=$(BUILD_NUMBER) -X main.ref=$(REF) -s -w" -o release/$(cmd)-linux-386 $(exe)

release-mac-amd64:
	$(info INFO: Starting build $@)
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags "-X main.version=$(VERSION) -X main.build=$(BUILD_NUMBER) -X main.ref=$(REF) -s -w" -o release/$(cmd)-mac-amd64 $(exe)

release-windows-amd64:
	$(info INFO: Starting build $@)
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags "-X main.version=$(VERSION) -X main.build=$(BUILD_NUMBER) -X main.ref=$(REF) -s -w" -o release/$(cmd)-windows-amd64.exe $(exe)

release-windows-386:
	$(info INFO: Starting build $@)
	CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -ldflags "-X main.version=$(VERSION) -X main.build=$(BUILD_NUMBER) -X main.ref=$(REF) -s -w" -o release/$(cmd)-windows-386.exe $(exe)

docker-image:
	$(info INFO: Starting build $@)
	docker build --build-arg VERSION --build-arg BUILD_NUMBER --build-arg REF -t simonbaeumer/monorepo-operator:$(VERSION) .

docker-push:
	$(info INFO: Starting build $@)
	docker push simonbaeumer/monorepo-operator:$(VERSION)

release: release-amd64 release-arm release-386 release-mac-amd64 release-windows-386 release-windows-amd64