# BUILD STAGE
FROM golang:1.12

ENV GO111MODULE on

ARG VERSION="0.0.0"
ARG BUILD_NUMBER="dev"
ARG REF="dev"

ENV VERSION ${VERSION}
ENV BUILD_NUMBER ${BUILD_NUMBER}
ENV REF ${REF}

WORKDIR /go/src/github.com/SimonBaeumer/monorepo-operator
COPY . ./
RUN make release-amd64

# IMAGE STAGE
FROM ubuntu:18.04

COPY --from=0 /go/src/github.com/SimonBaeumer/monorepo-operator/release/monorepo-operator-linux-amd64 /bin/monorepo-operator
RUN chmod +x /bin/monorepo-operator

RUN apt-get update
RUN apt-get install -y \
        wget \
        git
