start:
	go run cmd/main.go

test:
	GIN_MODE=release go test tests/main_test.go -v

start-test:
	make test && make start