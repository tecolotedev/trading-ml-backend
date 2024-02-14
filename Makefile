BINARY_NAME=my_app
DB_NAME=app_db
DB_USERNAME=admin
DB_PASSWORD=secret

## dockerDB: Create docker image for postgres
initPostgres:
	@echo "Initiating postgres..."
	docker run --name postgres15 -p 5432:5432 -e POSTGRES_USER=${DB_USERNAME} -e POSTGRES_PASSWORD=${DB_PASSWORD} -d postgres:15.0-alpine
	@echo "postgres initiated!!!"

## dropPostgres: it removes db from docker image
createPostgres:
	@echo "Creating postgres db..."
	docker exec -it postgres15 createdb --username=${DB_USERNAME} --owner=${DB_USERNAME} ${DB_NAME} 
	@echo "DB created!!!"

## dropPostgres: it removes db from docker image
dropPostgres:
	@echo "Dropping postgres db..."
	docker exec -it postgres15 dropdb --username=${DB_USERNAME} ${DB_NAME} 
	@echo "DB dropped!!!"

## migratePosgres: update tables inside database
migratePostgres:
	@echo "Migrating postgres db..."
	migrate -path sqlc/migrations -database "postgresql://${DB_USERNAME}:${DB_PASSWORD}@127.0.0.1:5432/${DB_NAME}?sslmode=disable" --verbose up
	@echo "DB migrated!!!"

## sqlc: Generate sqlc code
sqlcGenerate:
	@echo "Generating sqlc code..."
	sqlc generate
	@echo "sqlc code geneted!!!"

## build: Build binary
build:
	@echo "Building..."
	env CGO_ENABLED=0  go build -ldflags="-s -w" -o ${BINARY_NAME} .
	@echo "Built!"

## run: builds and runs the application
run: build
	@echo "Starting..."
	./${BINARY_NAME} &
	@echo "Started!"

## clean: runs go clean and deletes binaries
clean:
	@echo "Cleaning..."
	@go clean
	@rm ${BINARY_NAME}
	@echo "Cleaned!"

## start: an alias to run
start: run

## stop: stops the running application
stop:
	@echo "Stopping..."
	@-pkill -SIGTERM -f "./${BINARY_NAME}"
	@echo "Stopped!"

## restart: stops and starts the application
restart: stop start

## test: runs all tests
test:
	go test -v ./...