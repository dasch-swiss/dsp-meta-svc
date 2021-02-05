# Determine this makefile's path.
# Be sure to place this BEFORE `include` directives, if any.
# THIS_FILE := $(lastword $(MAKEFILE_LIST))
THIS_FILE := $(abspath $(lastword $(MAKEFILE_LIST)))
CURRENT_DIR := $(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))

include vars.mk

#################################
# Bazel targets
#################################

.PHONY: build
build: ## build all targets
	@yarn run build
	@bazel build //...

#################################
# Docker targets
#################################

.PHONY: docker-build
docker-build: # build and publish knora-api docker image locally
	@bazel run //src/server:app_image

.PHONY: docker-publish
docker-publish: # publish knora-api image to Dockerhub
	@bazel run //docker/knora-api:push
