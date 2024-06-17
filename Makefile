BROKER_BINARY=broker

# up: starts all containers in the background without forcing build
up:
	@echo "[codemore.io] Starting Docker images..."
	docker-compose up -d
	@echo "[codemore.io] Containers has successfully been started!"

# down: stopping docker containers
down:
	@echo "[codemore.io] Stopping containers..."
	docker-compose down
	@echo "[codemore.io] Containers has successfully been stopped!"

# up-build: stops docker-compose (if running), builds all projects and starts docker compose
up-build: build-broker
	@echo "[codemore.io] Stopping docker images (if running...)"
	docker-compose down
	@echo "[codemore.io] Building (when required) and starting docker images..."
	docker-compose up --build -d
	@echo "[codemore.io] Docker images built and started!"

# build-broker: 
build-broker:
	@echo "[codemore.io] Building broker..."
	cd ./broker && env GOOS=linux CGO_ENABLED=0 go build -o ${BROKER_BINARY} ./cmd/api
	@echo "[codemore.io] Broker has successfully been built!"

# users-migrate-up: run goose migrate up for users database
users-migrate-up:
	@echo "Running goose up migration on users database..."
	goose -dir ./broker/sql/migrations postgres postgresql://postgres:postgres@localhost:5432/users up
	@echo "Successfully ran goose up migration!"

# users-migrate-down: run goose migrate down for users database
users-migrate-down:
	@echo "Running goose down migration on users database..."
	goose -dir ./broker/sql/migrations postgres postgresql://postgres:postgres@localhost:5432/users down
	@echo "Successfully ran goose down migration!"

# client: start client
client:
	@echo "[codemore.io] Starting client..."
	cd ./client && npm run dev