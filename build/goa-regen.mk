# Goa "smart" regeneration
# Author: Raphaël Simon - @raphael
# Original: https://github.com/goadesign/examples/tree/master/code_regen
design_path="$(shell pwd | sed "s|^${GOPATH}/src/||")/design"

.PHONY: regen_goa
regen_goa: backup clean_goa gen_goa restore_goa

.PHONY: backup
backup:
	@echo "backing up"
	@find . -maxdepth 1 -name "*.go" -exec cp -f {} {}.backup \;

.PHONY: clean_goa
clean_goa:
	@echo "cleaning"
	@rm -rf app client swagger tool
	@find . -maxdepth 1 -name "*.go" -exec rm -f {} \;

.PHONY: gen_goa
gen_goa:
	@echo "generating $(design_path)"
	@goagen bootstrap -d $(design_path)

.PHONY: restore_goa
restore_goa:
	@echo "restoring"
	@build/goa-regen.py

.PHONY: clean_backups
clean_backups:
	@find . -maxdepth 1 -name "*.backup" -exec rm -f {} \;

.PHONY: clean_all
clean_all: clean_goa clean_backups
