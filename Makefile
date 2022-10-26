
LOCALPWD=$(shell pwd)
PACKAGE=$(shell basename $(LOCALPWD))

GOPATH=$(shell go env GOPATH)

.PHONY: dependence
build: ## make deps and checksum of direct/indirect dep 
	go mod init $(PACKAGE)
	go mod tidy


.PHONY: mkdirmocks
mkdirmocks:
	mkdir -p pkg/mocks

.PHONY: mockprepare
mockprepare: mkdirmocks
	go get github.com/golang/mock/gomock
	go get github.com/golang/mock/mockgen
	$(GOPATH)/bin/mockgen -source pkg/hello.go -destination pkg/mocks/mock_hello.go -package himocks

.PHONY: test
test: mockprepare mkdirmocks
	go test -v ./...

.PHONY: clean
clean: 
	rm -rf go.mod go.sum pkg/mocks
