all: deps compile

deps:
	@echo "installing dependencies..."
	@go get ./...

compile:
	@echo "compiling..."
	@go build ./...

format:
	@gofmt -s -w -l .
	@goimports -w .
