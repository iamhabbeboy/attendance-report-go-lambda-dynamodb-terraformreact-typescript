.PHONY: lint task build
lint: 
	golangci-lint run
run:
	air

build:
	env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w"  -o bin/main *.go

deploy: build
	sls deploy --stage prod

local:
	go build -o bin/main *.go && sls invoke local --function attendance