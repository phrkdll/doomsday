.PHONY: help

help: # Display this help screen
	@echo "Targets:"
	@echo
	@sed -n 's/^\([A-Za-z0-9_.-]*\):.*# \(.*\)$$/\1: \2/p' Makefile | sort | column -t -s ':'
	@echo

proto: # Generate Go code from .proto files
	@protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./protos/*.proto