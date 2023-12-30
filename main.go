package main

import (
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)
func main () {
	lambda.Start(handler)
}

// creating a lambda handler that handles apigateway requests and returns a response or an error
func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Println("hello")

	response := events.APIGatewayProxyResponse{
		StatusCode: 200,
	}
	return response, nil

}