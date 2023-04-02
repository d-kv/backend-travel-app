.SILENT: docs.generate
docs.generate:
	@cd docs && $(MAKE) --no-print-directory generate-uml

.SILENT: tools.download
tools.download:
	@cat tools/tools.go | grep _ | awk -F'"' '{print $$2}' | xargs -tI % go install %

.SILENT: tools.install
tools.install: tools.download
	@go mod download
	@cd docs && $(MAKE) --no-print-directory install-tools

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
