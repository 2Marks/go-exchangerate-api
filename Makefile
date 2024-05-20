build: 
	@go build -o bin/go-exchangerate-api cmd/main.go

test:
	@go test -v ./..

run: build
	@./bin/go-exchangerate-api