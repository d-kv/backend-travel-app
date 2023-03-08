.SILENT: generate-docs generate-uml
generate-docs:
	@cd docs && $(MAKE) --no-print-directory generate-uml

.SILENT: install-tools
install-tools:
	@cd docs && $(MAKE) --no-print-directory install-tools
