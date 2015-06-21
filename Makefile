 
# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOGET=$(GOCMD) get -t
GOINSTALL=$(GOCMD) install
GORUN=$(GOCMD) run
GOTEST=$(GOCMD) test
GODEP=$(GOTEST) -i
GOFMT=gofmt -w

PKG=github.com/msempere/aeropipe
PKG_JSON=gopkg.in/mgo.v2/bson
PKG_TERMUTIL=github.com/andrew-d/go-termutil
PKG_CLI=github.com/codegangsta/cli
PKG_AEROSPIKE=github.com/aerospike/aerospike-client-go

BIN=bin
USER_BIN=/usr/local/bin
APPTARGET=$(BIN)/aeropipe
APPSRC=aeropipe.go

all: clean build

build: dependencies
	$(GOBUILD) -o $(APPTARGET) $(APPSRC)

install:
	cp $(APPTARGET) $(USER_BIN)

dependencies:
	$(GOGET) $(PKG_JSON)
	$(GOGET) $(PKG_TERMUTIL)
	$(GOGET) $(PKG_CLI)
	$(GOGET) $(PKG_AEROSPIKE)
	$(GOGET) $(PKG)

run:
	./$(APPTARGET)

clean:
	rm -rf $(BIN)
	rm -rf $(USER_BIN)/aeropipe
