start:
	GIN_MODE=release go test tests/main_test.go -v && go run cmd/main.go