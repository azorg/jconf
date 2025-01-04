# File: "Makefile"

PRJ="github.com/azorg/jconf"

GIT_MESSAGE = "auto commit"

# go source files, ignore vendor directory
SRC = $(shell find . -type f -name '*.go' -not -path "./vendor/*")

# go packages
PKGS = $(PRJ)

.PHONY: all help clean distclean commit tidy vendor fmt test

all: fmt doc test

help:
	@echo "make fmt       - format Go sources"
	@echo "make test      - run test"
	@echo "make distclean - full clean (go.mod, go.sum)"

clean:
	@rm -f doc.txt
	@rm -f doc.md

distclean: clean
	@rm -f go.mod
	@rm -f go.sum
	@#sudo rm -rf go/pkg
	@rm -rf vendor
	@go clean -modcache
	
commit: fmt
	git add .
	git commit -am $(GIT_MESSAGE)
	git push

go.mod:
	@go mod init $(PRJ)
	@touch go.mod

tidy: go.mod
	@go mod tidy

go.sum: go.mod Makefile #tidy
	@#go get golang.org/x/exp/slog # experimental slog (go <1.21)
	@#go get github.com/azorg/xlog@go1.20
	@go get github.com/azorg/xlog
	@go get sigs.k8s.io/yaml
	@#go get github.com/ghodss/yaml
	@go get github.com/itchyny/json2yaml
	@go get github.com/stretchr/testify/require
	@touch go.sum

vendor: go.sum
	@go mod vendor

fmt: go.mod go.sum
	@go fmt

simplify:
	@gofmt -l -w -s $(SRC)

vet:
	@#go vet
	@go vet $(PKGS)

test: go.mod go.sum
	@go test

doc: doc.txt doc.md

doc.txt: *.go
	go doc -all > doc.txt

doc.md: *.go ~/go/bin/gomarkdoc
	~/go/bin/gomarkdoc -o doc.md

~/go/bin/gomarkdoc:
	go install github.com/princjef/gomarkdoc/cmd/gomarkdoc@latest

# EOF: "Makefile"
