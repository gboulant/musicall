pathname=$(shell go list)
basename=$(shell basename ${pathname})
PROGNAME=${basename}

all: build

build:
	@go build

test: build
	@./${PROGNAME}

clean:
	@go clean
	@rm -f *~ output.*
