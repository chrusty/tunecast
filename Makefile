default: api_codegen mocks build

api_codegen:
	@echo "Generating API from api/spec.yaml..."
	@oapi-codegen -package=api api/spec.yaml >api/api.gen.go

deps:
	@echo "Installing 'counterfeiter'..."
	@go get github.com/maxbrunsfeld/counterfeiter/v6@master

mocks:
	@echo "Creating Client mock (api/mock/client.go)..."
	@mkdir -p api/mock
	@counterfeiter -o api/mock/client.go api/api.gen.go ClientInterface

build:
	@echo " => Building (./bin) ..."
	@mkdir -p bin
	@echo "Building bin/tunecast ..."
	@go build -o bin/tunecast cmd/tunecast/main.go

test:
	@echo " => Running tests ..."
	@go test ./... -cover
