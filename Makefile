
SHELL := /bin/bash

init:
	chmod +x .pre-commit
	cp .pre-commit .git/hooks/pre-commit
