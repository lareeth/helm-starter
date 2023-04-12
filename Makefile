HELM_HOME ?= $(shell helm home)
HELM_PLUGINS := $(shell bash -c 'eval $$(helm env); echo $$HELM_PLUGINS')

VERSION := $(shell sed -n -e 's/version:[ "]*\([^"]*\).*/\1/p' plugin.yaml)

LDFLAGS := "-X main.version=${VERSION}"

.PHONY: bootstrap
bootstrap:
	go mod download

.PHONY: build
build:
	mkdir -p bin/
	go build -o bin/starter -ldflags $(LDFLAGS) ./main.go

.PHONY: install
install: build
	mkdir -p $(HELM_PLUGINS)/helm-starter/bin
	cp bin/starter $(HELM_PLUGINS)/helm-starter/bin
	cp plugin.yaml $(HELM_PLUGINS)/helm-starter/

.PHONY: dist
dist:
	mkdir -p dist/{linux,darwin,windows}

	GOOS=linux GOARCH=amd64 go build -o dist/linux/bin/starter -ldflags $(LDFLAGS) ./main.go
	tar -C dist/linux -zcvf dist/helm-starter-linux-$(VERSION).tgz bin/starter ../../plugin.yaml

	GOOS=darwin GOARCH=amd64 go build -o dist/darwin/bin/starter -ldflags $(LDFLAGS) ./main.go
	tar -C dist/darwin -zcvf dist/helm-starter-macos-$(VERSION).tgz bin/starter ../../plugin.yaml

	GOOS=windows GOARCH=amd64 go build -o dist/windows/bin/starter.exe -ldflags $(LDFLAGS) ./main.go
	tar -C dist/windows -zcvf dist/helm-starter-windows-$(VERSION).tgz bin/starter.exe ../../plugin.yaml