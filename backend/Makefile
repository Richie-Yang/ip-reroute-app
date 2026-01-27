-include env/.env.deployment
export

# Variables (env vars override these defaults)
PROJECT_ID ?= PROJECT_ID
SERVICE_NAME ?= SERVICE_NAME
REGION ?= REGION

# Development
.PHONY: dev
dev:
	go run main.go

.PHONY: build
build:
	go build -o bin/main .

# Docker
.PHONY: docker-build
docker-build:
	docker build -t $(SERVICE_NAME) .

.PHONY: docker-run
docker-run:
	docker run -p 8080:8080 $(SERVICE_NAME)

# Cloud Run Deployment
.PHONY: deploy
deploy:
	gcloud run deploy $(SERVICE_NAME) \
		--project $(PROJECT_ID) \
		--source . \
		--region $(REGION) \
		--set-env-vars "GIN_MODE=release" \
		--set-env-vars "ENV_MODE=prod" \
		--allow-unauthenticated

.PHONY: logs
logs:
	gcloud run services logs read $(SERVICE_NAME) --region $(REGION)

# Utilities
.PHONY: tidy
tidy:
	go mod tidy

.PHONY: clean
clean:
	rm -rf bin/