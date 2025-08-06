.PHONY: help

help: # Display this help screen
	@echo "Targets:"
	@echo
	@sed -n 's/^\([A-Za-z0-9_.-]*\):.*# \(.*\)$$/\1: \2/p' Makefile | sort | column -t -s ':'
	@echo

watch: # Perform "templ generate --watch" with additional options
	@templ generate --watch --proxy="http://localhost:8081" --cmd="go run ." --open-browser=false
