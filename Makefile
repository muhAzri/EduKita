.PHONY: build-local
build-local: ## Build the local docker image.
	docker compose -f docker/local/docker-compose.yml build

.PHONY: start-local
start-local: ## Start the local docker container.
	docker compose -f docker/local/docker-compose.yml up -d
