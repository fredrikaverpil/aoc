.PHONY: run
run: ## Run all solutions
	go run main.go

.PHONY: test
test: ## Run all tests
	go test -v ./...

.PHONY: test-day
test-day: ## Run tests for a specific day (usage: make test-day DAY=01)
	@if [ -z "$(DAY)" ]; then \
		echo "Error: DAY is required. Usage: make test-day DAY=01"; \
		exit 1; \
	fi
	go test -v ./day$(DAY)

.PHONY: bench
bench: ## Run all benchmarks
	go test -bench=. -benchmem ./...

.PHONY: bench-day
bench-day: ## Run benchmarks for a specific day (usage: make bench-day DAY=01)
	@if [ -z "$(DAY)" ]; then \
		echo "Error: DAY is required. Usage: make bench-day DAY=01"; \
		exit 1; \
	fi
	go test -bench=. -benchmem ./day$(DAY)
