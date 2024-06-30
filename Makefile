.PHONY: .test
.test:
	$(info Running tests...)
	go test ./...

.PHONY: test
test: .test ## run unit tests