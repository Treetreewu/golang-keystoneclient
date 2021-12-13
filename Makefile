test-unit:
	echo "Empty unit test."

.bin/golangci-lint: Makefile
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b .bin v1.39.0

test-style:
	golangci-lint run -v ./...

coverage:
	make test-unit
