CURRENT_DIR := $(shell pwd)

proto-gen:
	./scripts/gen-proto.sh ${CURRENT_DIR}

swag-init:
	swag init -g api/router.go --output api/docs