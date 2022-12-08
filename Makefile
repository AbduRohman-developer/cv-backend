CURRENT_DIR=$(shell pwd)
APP=$(shell basename ${CURRENT_DIR})

APP_CMD_DIR=${CURRENT_DIR}/cmd
PKG_LIST := $(shell go list ./... | grep -v /vendor/)

ifneq (,$(wildcard ./.env))
	include .env
endif

make create-env:
	cp ./.env.example ./.env

set-env:
	${CURRENT_DIR}/scripts/set-env.sh

build:
	CGO_ENABLED=0 GOOS=linux go build -mod=vendor -a -installsuffix cgo -o ${CURRENT_DIR}/bin/${APP} ${APP_CMD_DIR}/main.go

run:
	sudo go run cmd/main.go

clear:
	rm -rf ${CURRENT_DIR}/bin/*

run-dev: set-env
	sudo go run cmd/main.go

dep:
	go get -v -d ./...

lint:
	golint -set_exit_status ${PKG_LIST}

unit-tests: set-env ## Run unit-tests
	go test -mod=vendor -v -cover -short ${PKG_LIST}

race: set-env ## Run data race detector
	go test -mod=vendor -race -short ${PKG_LIST}

delete-branches:
	${CURRENT_DIR}/scripts/delete-branches.sh
vendor:
	go mod vendor -v

