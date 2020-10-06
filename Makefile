# structure from argoproj

PACKAGE=github.com/jw3/ppc
CURRENT_DIR=$(shell pwd)
DIST_DIR=${CURRENT_DIR}/dist

BUILD_DATE=$(shell date -u +'%Y-%m-%dT%H:%M:%SZ')
GIT_COMMIT=$(shell git rev-parse HEAD)
GIT_TAG=$(shell if [ -z "`git status --porcelain`" ]; then git describe --exact-match --tags HEAD 2>/dev/null; fi)
GIT_TREE_STATE=$(shell if [ -z "`git status --porcelain`" ]; then echo "clean" ; else echo "dirty"; fi)
VERSION=?${GIT_TAG}

FW_BUILD_DIR=${DIST_DIR}/firmware
export FW_SRC_DIR=${CURRENT_DIR}/firmware
export FW_VERSION=${VERSION}

override LDFLAGS += \
  -X ${PACKAGE}.version=${VERSION} \
  -X ${PACKAGE}.buildDate=${BUILD_DATE} \
  -X ${PACKAGE}.gitCommit=${GIT_COMMIT} \
  -X ${PACKAGE}.gitTreeState=${GIT_TREE_STATE}

#  docker image publishing options
DOCKER_PUSH?=false
IMAGE_NAMESPACE?=jwiii
IMAGE_TAG?=latest
GOARCH?=amd64
CGO_ENABLED=0
GOOS=linux

ifeq (${DOCKER_PUSH},true)
ifndef IMAGE_NAMESPACE
$(error IMAGE_NAMESPACE must be set to push images (e.g. IMAGE_NAMESPACE=jwiii))
endif
endif

ifneq (${GIT_TAG},)
IMAGE_TAG=${GIT_TAG}
override LDFLAGS += -X ${PACKAGE}.gitTag=${GIT_TAG}
endif

ifdef IMAGE_NAMESPACE
IMAGE_PREFIX=${IMAGE_NAMESPACE}/
endif

ifeq (${GOARCH},arm)
export GOARM=7
endif

# Build the project images
.DELETE_ON_ERROR:
all: ppc cli

all-images: ppc-image

.PHONY: all server-config ppc ppc-image firmware cli

# private cloud server
server-config:
	go build -v -ldflags '${LDFLAGS}' ./servers/ServerConfig.go

ppc: server-config
	go build -v -ldflags '${LDFLAGS}' -o ${DIST_DIR}/ppc ./servers/cmd/main.go

ppc-image:
	docker build --build-arg https_proxy -t $(IMAGE_PREFIX)ppc:$(IMAGE_TAG) -f ./servers/cmd/Dockerfile .
	@if [ "$(DOCKER_PUSH)" = "true" ] ; then  docker push $(IMAGE_PREFIX)ppc:$(IMAGE_TAG) ; fi

firmware:
	conan export-pkg $(FW_SRC_DIR) "jw3/stable" -s "compiler.version=5" -sf ${FW_BUILD_DIR} -f

cli:
	go build -v -ldflags '${LDFLAGS}' -o ${DIST_DIR}/polyform ./cli/cmd/main.go
