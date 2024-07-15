CURRENT_DIR := $(shell pwd)

proto-gen:
	./scripts/gen-proto.sh ${CURRENT_DIR}