.SILENT: generate-docs generate-uml
generate-docs:
	@cd docs && $(MAKE) --no-print-directory generate-uml

.SILENT: install-tools
install-tools:
	@cd docs && $(MAKE) --no-print-directory install-tools

.SILENT: tools.download
tools.download:
	@cat tools/tools.go | grep _ | awk -F'"' '{print $$2}' | xargs -tI % go install %


.SILENT: tools.install
tools.install: tools.download
	@go mod download

docker-compose.up:
	@(cd deployment && docker-compose up -d)

docker-compose.down:
	@(cd deployment && docker-compose down)
