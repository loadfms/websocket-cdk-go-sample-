package main

import (
	"cdk-go-sample/internal/cloud/lambda"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type CdkGoSampleStackProps struct {
	awscdk.StackProps
}

func NewCdkGoSampleStack(scope constructs.Construct, id string, props *CdkGoSampleStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	//SOCKET API GATEWAY
	socketApi := awsapigatewayv2.NewWebSocketApi(stack, jsii.String("HelloWebSocketApi"), &awsapigatewayv2.WebSocketApiProps{
		RouteSelectionExpression: jsii.String("$request.body.action"),
	})

	awsapigatewayv2.NewWebSocketStage(stack, jsii.String("HelloWebSocketStage"), &awsapigatewayv2.WebSocketStageProps{
		WebSocketApi: socketApi,
		StageName:    jsii.String("dev"),
		AutoDeploy:   jsii.Bool(true),
	})

	//ROUTES
	lambda.RegisterHelloRoutes(stack, socketApi)

	return stack
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	NewCdkGoSampleStack(app, "CdkGoSampleStack", &CdkGoSampleStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

func env() *awscdk.Environment {
	return &awscdk.Environment{
		Region: jsii.String("sa-east-1"),
	}
}
