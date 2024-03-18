.PHONY: run-rest mock

run:
	go run ./cmd/api/

test:
	go test -coverpkg=./... -race -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out | grep total
	rm coverage.out
