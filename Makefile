#!/bin/bash

TARGET=demo.a

GOPATH_DIR=/home/orcno/goloang-study
GCC=go
GBUILD=$(GCC) build -v
MODULE:=hannuo sort runetest slice structtest gutils 

.PHONY:all clean install


all:build


build: 
	env GOPATH="$(GOPATH_DIR):${GOPATH}" $(GBUILD) -o $(TARGET)


${MODULE}:
	env GOPATH="$(GOPATH_DIR):${GOPATH}" $(GBUILD) -o $@.o $@/


install:



clean:
	rm -f $(addsuffix .o,${MODULE})
	rm -f ${TARGET}

