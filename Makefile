OUTPUT_DIR=./bin
BIN_NAME=httpbin2
IMAGE_REPO=ghcr.io/trstringer/httpbin2

.PHONY: build
build:
	mkdir -p $(OUTPUT_DIR)
	go build -o $(OUTPUT_DIR)/$(BIN_NAME) .

.PHONY: image-build
image-build: build
	VERSION=$$($(OUTPUT_DIR)/$(BIN_NAME) version); \
	docker build -t $(IMAGE_REPO):$$VERSION -t $(IMAGE_REPO):latest .

.PHONY: image-push
image-push: build
	VERSION=$$($(OUTPUT_DIR)/$(BIN_NAME) version); \
	docker push $(IMAGE_REPO):$$VERSION; \
	docker push $(IMAGE_REPO):latest
