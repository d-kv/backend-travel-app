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

.SILENT: docker.build
docker.build:
	@docker build \
		--file Dockerfile \
		--tag afterwork \
		.

.SILENT: docker-compose.up
docker-compose.up: docker.build
	@(cd deployment && docker-compose up -d)

.SILENT: docker-compose.down
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

.SILENT: lint
lint: _lint_vet _lint_imports _lint_golangci

.SILENT: short-test.run
short-test.run:
	@go test ./... -count=1 -v -short

.SILENT: test.run
test.run:
	@go test ./... -count=1 -v

.SILENT: build
build:
	@go build cmd/place-service/main.go

.SILENT: mock.gen
mock.gen:
	@mockgen -source=pkg/adapter/igateway/oauth_provider.go -destination=internal/adapter/gateway/oauth_provider/mock/oauth_provider.go
	@mockgen -source=pkg/adapter/igateway/place_provider.go -destination=internal/adapter/gateway/place_provider/mock/place_provider.go
	@mockgen -source=pkg/app/icontroller/v0/controller.go -destination=internal/app/controller/v0/mock/controller.go
	@mockgen -source=pkg/infra/irepository/place.go -destination=internal/infra/repository/mock/place.go
	@mockgen -source=pkg/infra/irepository/user.go -destination=internal/infra/repository/mock/user.go
