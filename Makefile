.PHONY: clean
clean:
	rm -rf ./dist cover.out

.PHONY: test
test:
	go test \
		-race \
		-covermode atomic \
		-coverprofile=cover.out \
		./...
	go tool cover -func cover.out

.PHONY: code-lint
code-lint:
	golangci-lint run
