IMAGE_NAME = github.com/mgaliano18/discord-bot
SVC_NAME = discord-bot
APP_VSN := $(shell grep 'const Version = ' main.go | sed 's/^.*\s=\s//' | sed 's/\"//g')
BUILD := $(shell git rev-parse --short HEAD)

.PHONY: update
update:
	sh ./scripts/update_modules.sh
	make build

.PHONY: build
build:
	go mod vendor
	go build -mod=vendor -o $(SVC_NAME) ./cmd

.PHONY: check-format
#check-format:
#	gofmt -l ./pkg/... ./cmd/...
	
.PHONY: test
test: build
	go test -covermode=count -v -count=1  -coverpkg=$(COVERAGE_PKG) -coverprofile=coverage.out ./pkg/...
	go tool cover -func coverage.out | grep total

.PHONY: docker-build
docker-build:
	go mod vendor
	docker build --build-arg APP_VSN=$(APP_VSN) \
		--build-arg BUILD=$(BUILD) \
		-t $(IMAGE_NAME):$(APP_VSN)-$(BUILD) \
		-t $(IMAGE_NAME):latest .

.PHONY: docker-push
docker-push:
	docker push $(IMAGE_NAME):$(APP_VSN)-$(BUILD)
	docker push $(IMAGE_NAME):latest

.PHONY: linter
linter:
	revive -config linter.toml $(addprefix -exclude ,  $(GENFILES)) -formatter stylish ./

.PHONY: proto
proto:
	go mod vendor
	sh ./scripts/compile_proto.sh

.PHONY: easyjson
easyjson:
	sh ./scripts/generate_easyjson.sh
	go mod vendor
