GOPATH:= $(CURDIR)/vendor:$(GOPATH)

###############################################################################################
# code build
###############################################################################################
all: build

build:
	go build -v -o ./bin/sca .

generate:
	go mod vendor

clean:
	rm pkg/* -rf
	rm bin/* -f

