FROM --platform=$TARGETPLATFORM gcr.io/spectro-images-public/golang:1.21-alpine as builder
ARG TARGETOS
ARG TARGETARCH
ARG CRYPTO_LIB
ENV GOEXPERIMENT=${CRYPTO_LIB:+boringcrypto}

WORKDIR /workspace

COPY . .

RUN apk -U add coreutils gcc musl-dev

RUN mkdir -p bin

RUN if [ ${CRYPTO_LIB} ]; \
    then \
      GOARCH=${ARCH} go-build-fips.sh -a -o bin/cluster-autoscaler -mod=readonly main.go ;\
    else \
      GOARCH=${ARCH} go-build-static.sh -a -o bin/cluster-autoscaler -mod=readonly main.go ;\
    fi
RUN if [ "${CRYPTO_LIB}" ]; then assert-static.sh bin/cluster-autoscaler; fi
RUN if [ "${CRYPTO_LIB}" ]; then assert-fips.sh bin/cluster-autoscaler; fi
RUN scan-govulncheck.sh bin/cluster-autoscaler

FROM --platform=$TARGETPLATFORM gcr.io/distroless/static:nonroot-amd64
WORKDIR /
COPY --from=builder /workspace/bin/cluster-autoscaler .
CMD ["/cluster-autoscaler"]