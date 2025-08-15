all: build

build:
	@go build

test: build
	@./wavechart

clean:
	@go clean
	@rm -f *~
	@rm -rf output.*