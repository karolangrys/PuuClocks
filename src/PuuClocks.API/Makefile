build:
	go build ./...
run: 
	go run ./cmd/main.go
setup:
	go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest


lint:
	golangci-lint run ./... --timeout 10m

lint-fix:
	golangci-lint run ./... --fix

mocks:
	go generate ./...

test:
	ginkgo ./...

test-report:
	ginkgo --junit-report=report.xml  ./... 

swagger:
	swagger generate server --target=./internal/gen

openapi:
	make setup 
	oapi-codegen -generate gin,types -package gen_openapi ./api/openapi.yaml > ./gen/gen_openapi.go 
	

validate-openapi:
	go run github.com/getkin/kin-openapi/cmd/validate@latest -- ./api/openapi.yml 




