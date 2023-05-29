

SHELL := /bin/bash

.DEFAULT_GOAL := help

.PHONY: help
help: ## shows this help message
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'


.PHONY: up
up: ## docker-compose up
	docker-compose -f ./deployments/docker-compose.yaml  up --build


.PHONY: up-d
up-d: ## docker-compose up -d
	docker-compose -f ./deployments/docker-compose.yaml  up --build -d

.PHONY: down
down: ## docker-compose down
	docker-compose -f ./deployments/docker-compose.yaml  down

.PHONY: lint
lint: ## runs linter
	docker run  --rm -v "`pwd`:/app:cached" -w "/app/." golangci/golangci-lint:latest golangci-lint run