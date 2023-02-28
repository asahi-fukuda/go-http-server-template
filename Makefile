.PHONY: clean

OPENAPI_SPEC=doc/api.yaml

ts-stub: $(OPENAPI_SPEC)
	docker-compose run --rm gencode

go-server-stub:
	docker-compose run --rm app sh -c 'go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.10.1 && oapi-codegen -generate "types,server,spec" -package oapistub ./doc/api.yaml | tee ./pkg/oapistub/api.gen.go'

redoc: $(OPENAPI_SPEC)
	docker-compose run --rm redoc

all: ts-stub go-server-stub redoc

build:
	docker compose build

serve:
	docker compose up app
