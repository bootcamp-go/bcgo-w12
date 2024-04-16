.PHONY: test
test:
	go test -v ./...

.PHONY: test-coverage
test-coverage:
	go test -v -cover ./...


FILENAME=coverage.out

.PHONY: test-report
test-report:
	go test -v -cover -coverprofile=${FILENAME} ./...