// Copyright elipZis GmbH 2022
// All Rights Reserved
package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	fiberAdapter "github.com/awslabs/aws-lambda-go-api-proxy/fiber"
	"github.com/elipzis/go-serverless/web/router"
	_ "github.com/joho/godotenv/autoload"
	"log"
	"os"
	"os/signal"
)

var fiberLambda *fiberAdapter.FiberLambda
var r *router.Router

// init sets up some general routes and requirements per environment
func init() {
	// Init router
	r = router.NewRouter()
	r.Register(r.App.Group(""))

	// Fo AWS Lambda we need a wrapper to proxy the requests
	if os.Getenv("SERVER_ENV") == "AWS" {
		fiberLambda = fiberAdapter.New(r.App)
	}
}

// Handler to proxy AWS Lambda requests/responses
func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// If no name is provided in the HTTP request body, throw an error
	return fiberLambda.ProxyWithContext(ctx, req)
}

// Run, 8, 1
// @title Go Serverless! Example App
// @version 1.0
// @description The Example App for the Conf42: Golang 2022 Presentation "Go Serverless!"
// @contact.name Savas Ziplies
// @contact.web elipZis.com
// @contact.email goserverless@elipZis.com
// @BasePath /web
func main() {
	env := os.Getenv("SERVER_ENV")
	log.Printf("Starting on %s environment", env)
	if env == "AWS" {
		lambda.Start(Handler)
	} else {
		r.Run()
		quit := make(chan os.Signal)
		signal.Notify(quit, os.Interrupt)
		<-quit
	}
}
