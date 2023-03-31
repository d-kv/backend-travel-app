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

_lint_vet:
	@(cd cmd && go vet ./...)
	@(cd pkg && go vet ./...)
	@(cd internal && go vet ./...)

_lint_imports:
	@goimports-reviser cmd pkg internal tools

_lint_golangci:
	@golangci-lint run

lint: _lint_vet _lint_imports _lint_golangci
