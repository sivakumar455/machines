.PHONY: build test run clean

APP_NAME=machines
VERSION=0.0.1
PROJECT=machines
WORK_DIR=.

build:
	@echo "Building $(APP_NAME) version $(VERSION)..."
	go build -o ${APP_NAME} ${WORK_DIR}/cmd

test:
	go test ${WORK_DIR}/...

run: build
	./machines

clean:
	@echo "Deleting $(APP_NAME) version $(VERSION)..."
	rm -rf ${APP_NAME}