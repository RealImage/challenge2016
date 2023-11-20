DISTRIBUTION_BINARY = distributionApp

## build_front: builds the front end binary
build-app:
	@echo "Building front end binary..."
	env CGO_ENABLED=0 go build -o ./bin/${DISTRIBUTION_BINARY} ./cmd/api
	@echo "Done!"

## start: starts the front end
start: build-app
	@echo "Starting front end"
	go build -o ./bin/${DISTRIBUTION_BINARY} ./cmd/api
	./bin/${DISTRIBUTION_BINARY}