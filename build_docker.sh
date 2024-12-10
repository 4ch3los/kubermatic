#!/bin/bash
export DOCKER_REPO=registry.sva-cloud.com/kubermatic
export TAGS=v2.26.0
make docker-build
make docker-push