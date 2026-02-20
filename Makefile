#!/usr/bin/make -f

install:
	go install .

.PHONY: test fmt install