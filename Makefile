.SILENT: generate-uml
generate-docs:
	@cd docs && $(MAKE) generate-uml

install-tools:
	@cd docs && $(MAKE) install-tools
