.PHONY: build synth diff deploy zip clean fmt lint destroy

start:
	go run ./cmd/serve/main.go

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
	
build: 
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o lambdaHandler cmd/lambda/main.go && cd ..

zip:
	zip -Xr lambdaHandler.zip lambdaHandler resources

clean:
	rm lambdaHandler.zip lambdaHandler

deploy: lint fmt build zip cdkdeploy clean

cdkdeploy:
	cdk deploy --profile dev

destroy:
	cdk destroy --profile dev
