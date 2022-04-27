#BINARY=engine
#engine:
#	go build -o ${BINARY} cmd/*.go
test:
	go test ./...

container:
	docker build -t workshop-app .

run:
	docker-compose up --build -d

stop:
	docker-compose down

lint-prepare:
	@echo "Installing golangci-lint"
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s latest

lint:
	./bin/golangci-lint run ./...

.PHONY: clean install test build docker run stop vendor lint-prepare lint