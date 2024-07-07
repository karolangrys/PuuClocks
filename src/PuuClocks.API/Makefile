lint:
	golangci-lint run ./... --timeout 10m

lint-fix:
	golangci-lint run ./... --fix

mocks:
	go generate ./...

test:
	ginkgo ./...
