# Determine this makefile's path.
# Be sure to place this BEFORE `include` directives, if any.
# THIS_FILE := $(lastword $(MAKEFILE_LIST))
THIS_FILE := $(abspath $(lastword $(MAKEFILE_LIST)))
CURRENT_DIR := $(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))

include vars.mk

#################################
# Bazel targets
#################################

.PHONY: yarn
yarn: ## install dependencies
	@bazel run @nodejs//:yarn

.PHONY: build
build: yarn ## build all targets
	@bazel run @nodejs//:yarn -- run build
	@bazel build //...

.PHONY: test
test: yarn ## test all targets
	@bazel run @nodejs//:yarn -- run build
	@bazel test //...

.PHONY: buildifier
buildifier: ## format Bazel WORKSPACE and BUILD.bazel files
	@bazel run :buildifier

.PHONY: gen-go-deps
gen-go-deps: ## regenerate dependencies file (deps.bzl)
	@bazel run //:gazelle -- update-repos -from_file=go.mod -to_macro=deps.bzl%go_dependencies

.PHONY: docker-publish
docker-publish: metadata-docker-publish ## publish all docker images

#################################
# Metadata service targets
#################################

.PHONY: metadata
metadata: ## start Go mock backend on port 3000
	@go run services/metadata/backend/fake-backend/fake-backend.go

.PHONY: metadata-docker-build
metadata-docker-build: build ## publish metadata mock-server linux/amd64 platform docker image locally (watching /data/*.json)
	@bazel run --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64 //services/metadata/backend/fake-backend:image -- --norun

.PHONY: metadata-docker-publish
metadata-docker-publish: build ## publish metadata mock-server linux/amd64 platform docker image to Dockerhub (watching /data/*.json)
	@bazel run --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64 //services/metadata/backend/fake-backend:push

.PHONY: metadata-docker-run
metadata-docker-run: metadata-docker-build ## build and run metadata mock-server linux/amd64 platform docker image (watching /data/*.json)
	@docker run --rm -p 3000:3000 bazel/services/metadata/backend/fake-backend:image

.PHONY: metadata-service-run
metadata-service-run: build ## start the metadata-service
	@bazel run //services/metadata/backend/cmd

.PHONY: metadata-service-test
metadata-service-test: ## run all metadata-service tests
	@bazel test //services/metadata/backend/...

#################################
# Resource service targets
#################################

.PHONY: resource-docker-build
resource-docker-build: build ## publish linux/amd64 platform image locally
	@bazel run --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64 //services/resource/backend/cmd:image -- --norun

.PHONY: resource-docker-publish
resource-docker-publish: build ## publish linux/amd64 platform image to Dockerhub
	@bazel run --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64 //services/resource/docker:push

.PHONY: resource-service-run
resource-service-run: build ## start the resource-service
	@bazel run //services/resource/backend/cmd

.PHONY: resource-service-test
resource-service-test: ## run all resource-service tests
	@bazel test //services/resource/backend/...

#################################
# Other targets
#################################

.PHONY: help
help: ## this help
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST) | sort

.DEFAULT_GOAL := help
