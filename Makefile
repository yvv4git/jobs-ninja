lint:
	goimports -w .
	golangci-lint run -v --fix -c .lint-config.yml

run_collector_dump_history:
	go run cmd/collector/main.go history