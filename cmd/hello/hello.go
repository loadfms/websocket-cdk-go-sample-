// lambda/chat.go
package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/apigatewaymanagementapi"
)

type Response struct {
	Status int `json:"statusCode"`
}

func handler(ctx context.Context, event events.APIGatewayWebsocketProxyRequest) (Response, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		fmt.Println("Error loading AWS config:", err)
		return Response{Status: 500}, err
	}

	client := apigatewaymanagementapi.NewFromConfig(cfg, func(o *apigatewaymanagementapi.Options) {
		o.BaseEndpoint = aws.String("GET_WEBHOOK_URL_FROM_APIGATEWAY")
	})

	params := &apigatewaymanagementapi.PostToConnectionInput{
		ConnectionId: &event.RequestContext.ConnectionID,
		Data:         []byte("Hello! " + event.Body + " - " + event.RequestContext.ConnectionID),
	}

	result, err := client.PostToConnection(context.TODO(), params)
	if err != nil {
		fmt.Println("Error posting to connection:", err)
		return Response{Status: 500}, err
	}

	fmt.Println(result)

	return Response{Status: 200}, nil
}

func main() {
	lambda.Start(handler)
}
