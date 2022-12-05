WGEN_GO := go

WGEN_MAIN := supervisor.go
WGEN_OBJ := wgen

PREFIX := .
DESTDIR := bin

WGEN_YAML_WORKLOAD := ./yaml/robot-shop-workload.yml
WGEN_YAML_APISPEC := ./yaml/robot-shop-api.yml

WGEN_ARGS := -w $(WGEN_YAML_WORKLOAD) -a $(WGEN_YAML_APISPEC) -d 5s

.SILENT: help
.PHONY: help # print help
help:
	grep '^.PHONY: .* #' $(firstword $(MAKEFILE_LIST)) |\
	sed 's/\.PHONY: \(.*\) # \(.*\)/\1 # \2/' |\
	awk 'BEGIN {FS = "#"}; {printf "%-20s %s\n", $$1, $$2}'

.PHONY: build # compile wgen
build:
	$(WGEN_GO) build -o $(PREFIX)/$(DESTDIR)/$(WGEN_OBJ)

.PHONY: clean # clean and remove wgen binary
clean:
	$(WGEN_GO) clean
	rm -f $(PREFIX)/$(DESTDIR)/$(WGEN_OBJ)

.PHONY: run # compile and run wgen
run:
	clear
	$(WGEN_GO) run $(WGEN_MAIN) $(WGEN_ARGS)

.PHONY: install # install wgen into ~/.local/bin
install:
	install -D $(PREFIX)/$(DESTDIR)/$(WGEN_OBJ) ${HOME}/.local/bin/$(WGEN_OBJ)

.PHONY: uninstall # unininstall wgen from ~/.local/bin
uninstall:
	rm -f ${HOME}/.local/bin/$(WGEN_OBJ)

.PHONY: deps # resolve dependencies
deps:
	$(WGEN_GO) mod tidy

.PHONY: format # format all
format:
	$(WGEN_GO) fmt ./...
