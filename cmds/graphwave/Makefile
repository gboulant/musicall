all: build

build: wave.so

wave.so: wave.go
	go build -o $@ -buildmode=c-shared $<

test: build
	@python3 waveplot.py

clean:
	@rm -f wave.so wave.h
	@rm -rf __pycache__
