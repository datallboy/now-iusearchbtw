BINARY_NAME=now-iusearchbtw

build:
	mkdir -p bin
	CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o bin/${BINARY_NAME} cmd/now-iusearchbtw/main.go
	npm run build

run:
	APP_ENV=production bin/${BINARY_NAME}

build_and_run: build run

clean:
	go clean
	rm -rf ./bin