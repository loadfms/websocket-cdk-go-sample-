package lambda

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2integrations"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"

	"github.com/aws/jsii-runtime-go"
)

func RegisterHelloRoutes(stack awscdk.Stack, socketApi awsapigatewayv2.WebSocketApi) {
	//LAMBDAS
	helloLambda := awslambda.NewFunction(stack, jsii.String("HelloLambda"), &awslambda.FunctionProps{
		Runtime:      awslambda.Runtime_PROVIDED_AL2023(),
		Handler:      jsii.String("bootstrap"),
		Code:         awslambda.Code_FromAsset(jsii.String("bin/hello/hello.zip"), nil),
		Architecture: awslambda.Architecture_ARM_64(),
	})

	//INTEGRATIONS
	helloIntegration := awsapigatewayv2integrations.NewWebSocketLambdaIntegration(jsii.String("HelloIntegration"), helloLambda, &awsapigatewayv2integrations.WebSocketLambdaIntegrationProps{})

	socketApi.AddRoute(jsii.String("hello"), &awsapigatewayv2.WebSocketRouteOptions{
		Integration: helloIntegration,
	})

	socketApi.GrantManageConnections(helloLambda)
}
