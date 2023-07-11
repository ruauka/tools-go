#test:
#	@go test -cover ./... -coverprofile cover.out
#	@echo "-------------------------------------------------------------------------------------"
#	@go tool cover -func cover.out
#	@echo "-------------------------------------------------------------------------------------"
test:
	@go test . -coverprofile=coverage.out
	@go tool cover -func coverage.out

lint:
	@golangci-lint run

check:
	@echo "Test starting..."
	@make test
	@echo "Test ok..."
	@echo "Lint starting..."
	@make lint
	@echo "Lint ok..."