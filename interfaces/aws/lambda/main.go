package main

import (
	"context"
	"log"
	"os"

	"github.com/aokuyama/circle_scheduler-api/interfaces/aws/lambda/secret"
	"github.com/aokuyama/circle_scheduler-api/packages/infrastructure/http/gin/route"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
)

var ginLambda *ginadapter.GinLambda

func init() {
	// stdout and stderr are sent to AWS CloudWatch Logs
	log.Printf("Gin cold start")
	r := gin.Default()
	route.Define(r)

	ginLambda = ginadapter.New(r)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	e := os.Getenv("ENV_PREFIX")

	for _, key := range []string{
		"DATABASE_URL",
		"JWT_SECRET_KEY",
		"KEY_PATH",
		"PEPPER_PATH",
		"PEPPER_PASSWORD",
	} {
		err := secret.SetEnvBySecretParam(key, e)
		if err != nil {
			panic(err)
		}
	}
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}
