BINARY_NAME=myapp
DSN="host=localhost port=5432 user=postgres password=password dbname=concurrency sslmode=disable"
REDIS="localhost:6379"  # Add your Redis DSN if needed

## build: Build binary
build:
	@printf "🚀 Building...\n"
	@env CGO_ENABLED=0 go build -ldflags="-s -w" -o ${BINARY_NAME} ./cmd/web
	@printf "✅ Built!\n"

## run: builds and runs the application
run: build
	@echo "🚀 Starting..."
	@env DSN=${DSN} REDIS=${REDIS} ./${BINARY_NAME} &
	@echo "🎉 Started!"

## clean: runs go clean and deletes binaries
clean:
	@echo "🧹 Cleaning..."
	@go clean
	@rm -f ${BINARY_NAME}
	@echo "🗑 Cleaned!"

## start: an alias to run
start: run

## stop: stops the running application
stop:
	@echo "🛑 Stopping..."
	@-pkill -SIGTERM -f "./${BINARY_NAME}"
	@docker exec -it finalproject-redis-1  redis-cli FLUSHALL
	@echo "🛑 Stopped!"

## restart: stops and starts the application
restart: stop start

## test: runs all tests
test:
	@echo "🧪 Running tests..."
	@go test -v ./...
	@echo "✅ Tests completed!"
## load env: loads .env into the shell
env:
	@export $(grep -v '^#' .env | xargs)
