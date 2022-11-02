.PHONY: build synth diff deploy zip clean fmt lint destroy

start:
	nodemon -e go,json,html --exec go run ./cmd/serve/main.go --signal SIGTERM

bootstrap:
	cdk bootstrap --profile dev

synth:
	cdk synth --profile dev

diff:
	cdk diff --profile dev

fmt:
	go fmt ./...

lint:
	go vet ./...

test:
	go test ./...

build: 
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o lambdaHandler cmd/lambda/main.go && cd ..

zip:
	zip -Xr lambdaHandler.zip lambdaHandler resources

deploy: lint fmt build zip
	cdk deploy --profile dev

destroy:
	cdk destroy --profile dev
