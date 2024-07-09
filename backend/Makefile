SHELL:=/bin/bash
VERSION ?= $(shell git describe --match 'v[0-9]*' --tags --always)

migrate_up=go run . migration --action=up
migrate_down=go run . migration --action=down

changelog_args=-o CHANGELOG.md -tag-filter-pattern '^v'

changelog_args=-o CHANGELOG.md -tag-filter-pattern '^v'

build_loc=./bin/app
build_args=-ldflags "-s -w -X 'github.com/Halalins/backend/internal/bootstrap.ServiceName=crocodile' -X 'github.com/Halalins/backend/internal/bootstrap.ServiceVersion=$(VERSION)'" -o ${build_loc} main.go

.PHONY: run
run:
	go run . server

.PHONY: run_worker
run_worker:
	go run . worker

.PHONY: build
build:
	go build $(build_args)

.PHONY: migrate_up
migrate_up:
	@if [ "$(version)" = "" ]; then\
    	$(migrate_up);\
	else\
		go run . migration --action=up-to --version=$(version);\
    fi

.PHONY: migrate_down
migrate_down:
	@if [ "$(version)" = "" ]; then\
    	$(migrate_down);\
	else\
		go run . migration --action=down-to --version=$(version);\
    fi


.PHONY: migrate_create
migrate_create:
	@if [ "$(name)" = "" ]; then\
    	echo 'migration file need name' ;\
	else\
		go run . migration --action=create --name=$(name);\
    fi

.PHONY: changelog
changelog:
ifdef version
	$(eval changelog_args=--next-tag $(version) $(changelog_args))
endif
	git-chglog $(changelog_args)

.PHONY: lint
lint:
	golangci-lint run --print-issued-lines=false --exclude-use-default=false --fix --timeout=3m