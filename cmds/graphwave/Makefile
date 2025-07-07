all: build

build: wavelib.so

wavelib.so: wavelib.go
	go build -o $@ -buildmode=c-shared $<

test: build
	@python3 tester.py

clean:
	@rm -f wavelib.so wavelib.h
	@rm -rf __pycache__
