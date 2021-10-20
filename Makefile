.PHONY: all build install

all: build install


build:
	@go build

install:
	@go install
