TEST_POSTGRES_URL="postgres://postgres:passwd@localhost:7001/hackathons-api_test?sslmode=disable"
GOGETV = go get -v
GOBUILD = go build
MIGRATE = PATH="$$PWD/bin:$$PATH" \
           bin/migrate -database $(TEST_POSTGRES_URL)
.PHONY: prototool/install
#prototool/install: export GO111MODULE := on
prototool/install: export GOBIN := $(PWD)/bin
prototool/install:
	$(GOGETV) github.com/uber/prototool/cmd/prototool@v1.10.0
	$(GOGETV) github.com/golang/protobuf/protoc-gen-go@v1.4.2
   	#$(GOGETV) github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway@v1.14.6
   	#$(GOGETV) github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger@v1.14.6
.PHONY: generate
generate: prototool/install
	PATH="$$PWD/bin:$$PATH" prototool generate
.PHONY: database/up
database/up:
	@docker-compose up -d && sleep 5
.PHONY: database/down
database/down:
	@docker-compose down
.PHONY: migrate/build
migrations/build:
	@$(GOBUILD) -pkgdir=vendor/ -tags 'postgres' -o bin/migrate \
    	github.com/golang-migrate/migrate/v4/cmd/migrate
.PHONY: migrate/test
migrate/test:
	@make database/up
	@make migrations/build
	@$(MIGRATE) -source="file:$$PWD/storage/dbschema" up
.PHONY: test
test:
	@make migrate/test
	go test ./... || true
	@make database/down