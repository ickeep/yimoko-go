GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)
CONFIG_PROTO_FILES=$(shell find config -name *.proto)
INTERNAL_PROTO_FILES=$(shell find */*/internal -name *.proto)
API_PROTO_FILES=$(shell find api -name *.proto)
WIRE_FILES=$(shell find */*/cmd/* -name wire_gen.go)

.PHONY: proto
# generate api proto
proto:
	protoc --proto_path=. \
	       --proto_path=./third_party \
 	       --go_out=paths=source_relative:. \
				 --go-errors_out=paths=source_relative:. \
 	       --validate_out=paths=source_relative,lang=go:. \
	       $(API_PROTO_FILES)