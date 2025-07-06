all: build

build: wave.so

wave.so: wave.go
	go build -o $@ -buildmode=c-shared $<

clean:
	@rm -f wave.so wave.h
