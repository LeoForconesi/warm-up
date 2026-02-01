test:
	go test ./... -v -cover

race-test:
	go test -race ./...
