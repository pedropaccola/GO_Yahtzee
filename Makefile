build:
	@go build -o bin/yahtzee

run: build
	@./bin/yahtzee
	
test:
	@go test -v ./...