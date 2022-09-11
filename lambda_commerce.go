package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type LambdaCommerceStackProps struct {
	awscdk.StackProps
}

func NewLambdaCommerceStack(scope constructs.Construct, id string, props *LambdaCommerceStackProps) awscdk.Stack {
	var sProps awscdk.StackProps
	if props != nil {
		sProps = props.StackProps
	}

	stack := awscdk.NewStack(scope, &id, &sProps)

	lambdaCommerce := awslambda.NewFunction(stack, jsii.String("lambdaCommerce"), &awslambda.FunctionProps{
		Code:    awslambda.NewAssetCode(jsii.String("./cmd/build"), nil),
		Handler: jsii.String("lambdaHandler"),
		Timeout: awscdk.Duration_Seconds(jsii.Number(300)),
		Runtime: awslambda.Runtime_GO_1_X(),
	})

	functionUrlOptions := &awslambda.FunctionUrlOptions{
		AuthType: awslambda.FunctionUrlAuthType_NONE,
	}

	fnUrl := lambdaCommerce.AddFunctionUrl(functionUrlOptions)

	cfnOutputProps := &awscdk.CfnOutputProps{
		Value: fnUrl.Url(),
	}

	awscdk.NewCfnOutput(stack, jsii.String("TheUrl"), cfnOutputProps)

	return stack
}

func main() {
	app := awscdk.NewApp(nil)

	NewLambdaCommerceStack(app, "LambdaCommerceStack", &LambdaCommerceStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

// env determines the AWS environment (account+region) in which our stack is to
// be deployed. For more information see: https://docs.aws.amazon.com/cdk/latest/guide/environments.html
func env() *awscdk.Environment {
	return nil
}
