test:
	go test ./... -cover

run:
	air & cd views && pnpm dev && fg
	
build:
	rm -rf bin/main
	go build -o bin/main main.go
