export ALLURE_OUTPUT_PATH=$(CURDIR)

test:
	@rm -r allure-results || exit 0
	@go test ./... -coverpkg=./... -cover -coverprofile cover.out
	@echo "-------------------------------------------------------------------------------------"
	@go tool cover -func cover.out
	@echo "-------------------------------------------------------------------------------------"

lint:
	@golangci-lint run

allure:
	@allure serve allure-results -h 127.0.0.1

check:
	@echo "Test starting..."
	@make test
	@echo "Test ok..."
	@echo "Lint starting..."
	@make lint
