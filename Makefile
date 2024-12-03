.PHONY: build clean deploy

GO_BUILD := env GOARCH=arm64 GOOS=linux go build -ldflags="-s -w"

build:
	export GO111MODULE=on
	export CGO_ENABLED=1

	${GO_BUILD} -o bin/hello/bootstrap cmd/hello/hello.go
	chmod +x bin/hello/bootstrap
	cd bin/hello && pwd && zip -r hello.zip .

clean:
	rm -rf ./bin

deploy: clean build
	cdk bootstrap
	cdk deploy
