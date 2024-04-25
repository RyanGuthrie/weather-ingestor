
.PHONY: test
test:
	go test -race ./...

.PHONY: run
run:
	go run cmd/weather-ingestor/main.go

.PHONY: vet
vet:
	go vet ./...
