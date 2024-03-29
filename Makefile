EXCLUDE_THIRD_PARTY=--exclude-path third_party/errors --exclude-path third_party/google --exclude-path third_party/openapi --exclude-path third_party/validate

setup:
	go mod vendor
	go install github.com/cespare/reflex@latest
	go install github.com/pressly/goose/v3/cmd/goose@latest
	go install go.uber.org/mock/mockgen@latest

test:
	go test ./...

api:
	buf generate ${EXCLUDE_THIRD_PARTY} --path api/v1

build-api:
	go build -v -o bin/app-api cmd/app-api/*.go

build-cli:
	go build -v -o bin/app-cli cmd/app-cli/*.go

start-api-dev:
	make api
	reflex -r "\.(go|yaml)" -s -- sh -c "make build-api && ./bin/app-api -config=./files/config/development.yaml"