ALL_ARCH = amd64 arm64
all: $(addprefix build-arch-,$(ALL_ARCH))

SPECTRO_VERSION ?= 4.1.0-dev
TAG ?= v1.25.0-spectro-${SPECTRO_VERSION}
FLAGS=
LDFLAGS?=-s
ENVVAR=CGO_ENABLED=0
GOOS?=linux
GOARCH?=$(shell go env GOARCH)
#REGISTRY?=staging-k8s.gcr.io
DOCKER_NETWORK?=default
ifdef BUILD_TAGS
  TAGS_FLAG=--tags ${BUILD_TAGS}
  PROVIDER=-${BUILD_TAGS}
  FOR_PROVIDER=" for ${BUILD_TAGS}"
else
  TAGS_FLAG=
  PROVIDER=
  FOR_PROVIDER=
endif
ifdef LDFLAGS
  LDFLAGS_FLAG=--ldflags "${LDFLAGS}"
else
  LDFLAGS_FLAG=
endif
ifdef DOCKER_RM
  RM_FLAG=--rm
else
  RM_FLAG=
endif
IMAGE=$(REGISTRY)/cluster-autoscaler$(PROVIDER)

export DOCKER_CLI_EXPERIMENTAL := enabled

GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)
GOPATH ?= $(shell go env GOPATH)
TARGETARCH ?= linux/amd64

FIPS_ENABLE ?= ""

IMG_PATH ?= "gcr.io/spectro-dev-public/jayesh"
IMG_TAG ?= "latest"
IMG_SERVICE_URL ?= ${IMG_PATH}/
CAS_IMG ?= ${IMG_SERVICE_URL}cluster-autoscaler:${IMG_TAG}

BUILD_ARGS = --build-arg CRYPTO_LIB=${FIPS_ENABLE} --build-arg BUILDER_GOLANG_VERSION=${BUILDER_GOLANG_VERSION}
RELEASE_LOC := release
ifeq ($(FIPS_ENABLE),yes)
  RELEASE_LOC := release-fips
endif

ALL_ARCH = amd64 arm64
REGISTRY ?= gcr.io/spectro-dev-public/$(USER)/${RELEASE_LOC}
# main controller
CORE_IMAGE_NAME ?= cluster-autoscaler
CORE_CONTROLLER_IMG ?= $(REGISTRY)/$(CORE_IMAGE_NAME)

.PHONY: docker-build-all ## Build all the architecture docker images
docker-build-all: $(addprefix docker-build-,$(ALL_ARCH))

docker-build-%: ## Build docker images for a given ARCH
	$(MAKE) ARCH=$* docker-build

.PHONY: docker-build
docker-build:  ## Build the docker image for cluster autoscaler
	docker buildx build --load --platform linux/${ARCH} ${BUILD_ARGS} --build-arg ARCH=$(ARCH) --build-arg LDFLAGS="$(LDFLAGS)" . -t $(CORE_CONTROLLER_IMG)-$(ARCH):$(TAG)
	@echo $(CORE_CONTROLLER_IMG)-$(ARCH):$(TAG)

docker:
	docker buildx build --platform ${TARGETARCH} --push . -t ${CAS_IMG} --build-arg CRYPTO_LIB=${FIPS_ENABLE} -f Dockerfile.amd64

build: build-arch-$(GOARCH)

build-arch-%: clean-arch-%
	$(ENVVAR) GOOS=$(GOOS) GOARCH=$* go build -o cluster-autoscaler-$* ${LDFLAGS_FLAG} ${TAGS_FLAG}

test-unit: clean build
	go test --test.short -race ./... ${TAGS_FLAG}

dev-release: dev-release-arch-$(GOARCH)

dev-release-arch-%: build-arch-% make-image-arch-% push-image-arch-%
	@echo "Release ${TAG}${FOR_PROVIDER}-$* completed"

make-image: make-image-arch-$(GOARCH)

make-image-arch-%:
ifdef BASEIMAGE
	docker build --pull --build-arg BASEIMAGE=${BASEIMAGE} \
		-t ${IMAGE}-$*:${TAG} \
		-f Dockerfile.$* .
else
	docker build --pull \
		-t ${IMAGE}-$*:${TAG} \
		-f Dockerfile.$* .
endif
	@echo "Image ${TAG}${FOR_PROVIDER}-$* completed"

.PHONY: docker-push
docker-push: ## Push the docker image
	docker push $(CORE_CONTROLLER_IMG)-$(ARCH):$(TAG)

.PHONY: docker-push-all ## Push all the architecture docker images
docker-push-all: $(addprefix docker-push-,$(ALL_ARCH))
	$(MAKE) docker-push-core-manifest

docker-push-%: ## Docker push
	$(MAKE) ARCH=$* docker-push

.PHONY: docker-push-core-manifest
docker-push-core-manifest: ## Push the fat manifest docker image.
	## Minimum docker version 18.06.0 is required for creating and pushing manifest images.
	$(MAKE) docker-push-manifest CONTROLLER_IMG=$(CORE_CONTROLLER_IMG) MANIFEST_FILE=$(CORE_MANIFEST_FILE)

.PHONY: docker-push-manifest
docker-push-manifest: ## Push the manifest image
	docker manifest create --amend $(CONTROLLER_IMG):$(TAG) $(shell echo $(ALL_ARCH) | sed -e "s~[^ ]*~$(CONTROLLER_IMG)\-&:$(TAG)~g")
	@for arch in $(ALL_ARCH); do docker manifest annotate --arch $${arch} ${CONTROLLER_IMG}:${TAG} ${CONTROLLER_IMG}-$${arch}:${TAG}; done
	docker manifest push --purge ${CONTROLLER_IMG}:${TAG}

push-image: push-image-arch-$(GOARCH)

push-image-arch-%:
	./push_image.sh ${IMAGE}-$*:${TAG}

push-manifest:
	docker manifest create ${IMAGE}:${TAG} \
	    $(addprefix $(REGISTRY)/cluster-autoscaler$(PROVIDER)-, $(addsuffix :$(TAG), $(ALL_ARCH)))
	docker manifest push --purge ${IMAGE}:${TAG}

execute-release: $(addprefix make-image-arch-,$(ALL_ARCH)) $(addprefix push-image-arch-,$(ALL_ARCH)) push-manifest
	@echo "Release ${TAG}${FOR_PROVIDER} completed"

clean: clean-arch-$(GOARCH)

clean-arch-%:
	rm -f cluster-autoscaler-$*

generate:
	go generate ./cloudprovider/aws

format:
	test -z "$$(find . -path ./vendor -prune -type f -o -name '*.go' -exec gofmt -s -d {} + | tee /dev/stderr)" || \
	test -z "$$(find . -path ./vendor -prune -type f -o -name '*.go' -exec gofmt -s -w {} + | tee /dev/stderr)"

docker-builder:
	docker build --network=${DOCKER_NETWORK} -t autoscaling-builder ../builder

build-in-docker: build-in-docker-arch-$(GOARCH)

build-in-docker-arch-%: clean-arch-% docker-builder
	docker run ${RM_FLAG} -v `pwd`:/gopath/src/k8s.io/autoscaler/cluster-autoscaler/:Z autoscaling-builder:latest \
		bash -c 'cd /gopath/src/k8s.io/autoscaler/cluster-autoscaler && BUILD_TAGS=${BUILD_TAGS} LDFLAGS="${LDFLAGS}" make build-arch-$*'

release: $(addprefix build-in-docker-arch-,$(ALL_ARCH)) execute-release
	@echo "Full in-docker release ${TAG}${FOR_PROVIDER} completed"

container: container-arch-$(GOARCH)

container-arch-%: build-in-docker-arch-% make-image-arch-%
	@echo "Full in-docker image ${TAG}${FOR_PROVIDER}-$* completed"

test-in-docker: clean docker-builder
	docker run ${RM_FLAG} -v `pwd`:/cluster-autoscaler/:Z autoscaling-builder:latest bash -c 'cd /cluster-autoscaler && go test -race ./... ${TAGS_FLAG}'

.PHONY: all docker build test-unit clean format execute-release dev-release docker-builder build-in-docker release generate push-image push-manifest
