build:
	@go build -o bin/xml2json ./cmd/xml2json/main.go

run: build
	@./bin/xml2json

clean:
	@rm -rf bin/

test:
	@go test ./... -v
