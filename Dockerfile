FROM golang:1.21-alpine

ENV GOLANG_PROTOBUF_VERSION=1.28.1

ARG PROTOC_VERSION="3.20.0"
# add dependency
RUN apk add g++ make curl protoc git
# sanity check to verify its correctly installed
RUN protoc --version
# install
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v${GOLANG_PROTOBUF_VERSION}
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

WORKDIR /build
COPY . ./