lint:
	goimports -w .
	golangci-lint run -v --fix -c .lint-config.yml