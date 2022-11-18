build:
	@go build -o bin/yahtzee

run: build
	@go run ./bin/yahtzee
	
test:
	@go test -v ./...