.PHONY: build synth diff deploy fmt lint destroy

build: 
	cd cmd && GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o build/lambdaHandler lambda/main.go && zip -X build/lambdaHandler.zip build/lambdaHandler && cd ..

start:
	go run ./cmd/serve/main.go

bootstrap:
	cdk bootstrap --profile dev

synth:
	cdk synth --profile dev

diff:
	cdk diff --profile dev

deploy: lint fmt build
	cdk deploy --profile dev

fmt:
	go fmt ./...

lint:
	go vet ./...

destroy:
	cdk destroy --profile dev
