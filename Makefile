.PHONY: build-local
build-local: ## Build the local docker image.
	docker compose -f docker/local/docker-compose.yml build

.PHONY: start-local
start-local: ## Start the local docker container.
	docker compose -f docker/local/docker-compose.yml up -d

.PHONY: build-prod
build-prod: ## Build the production docker image.
	docker compose -f docker/production/docker-compose.yml build

.PHONY: start-prod
start-prod: ## Start the production docker container.
	docker compose -f docker/production/docker-compose.yml up -d

.PHONY: stop-prod
stop-prod: ## Stop the production docker container.
	docker compose -f docker/production/docker-compose.yml down

.PHONY: clean-prod
clean-prod: ## Clean the production docker container.
	docker compose -f docker/production/docker-compose.yml down -v