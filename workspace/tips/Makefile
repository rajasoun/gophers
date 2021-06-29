APP :=  $(shell basename $(CURDIR))
VERSION := $(shell git describe --tags --always --dirty)
GOPATH := $(CURDIR)/Godeps/_workspace:$(GOPATH)
PATH := $(GOPATH)/bin:$(PATH)

bin/$(APP): bin
	go build -v -o $@ -ldflags "-X main.Version='${VERSION}'"

bin: clean
	mkdir -p bin 

clean:
	rm -rf bin 

watch:
	gotestsum --watch --format testname

test:
	go test -v ./... --cover

.PHONY: bin/$(APP) bin clean test
