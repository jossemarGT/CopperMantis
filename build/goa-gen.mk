# Goa "smart" regeneration
# Author: Raphaël Simon - @raphael
# Original: https://github.com/goadesign/examples/tree/master/code_regen

.PHONY: goa_regen
goa_regen: goa_backup goa_clean goa goa_restore

.PHONY: goa_backup
goa_backup:
	@find . -maxdepth 1 -name "*.go" -exec cp -f {} {}.backup \;

.PHONY: goa_clean
goa_clean:
	@rm -rf app client swagger tool schema models
	@find . -maxdepth 1 -name "*.go" -exec rm -f {} \;

.PHONY: goa
goa:
	@go run build/goa_gen.go

.PHONY: goa_restore
goa_restore:
	@build/goa-regen.py

.PHONY: goa_backup_clean
goa_backup_clean:
	@find . -maxdepth 1 -name "*.backup" -exec rm -f {} \;
