package util

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"os"
)

// NewAWSSession returns a convenient session to connect to the services
func NewAWSSession() (sess *session.Session) {
	//If a custom url is configured (e.g. locally)
	if os.Getenv("ENVIRONMENT") == "local" {
		conf := aws.NewConfig().
			WithEndpoint(os.Getenv("QUEUE_URL")).
			WithRegion(os.Getenv("AWS_REGION")).
			WithCredentials(credentials.NewStaticCredentials("123", "123", ""))
		// Print Debugs
		if os.Getenv("DEBUG") == "true" {
			conf = conf.WithLogLevel(*aws.LogLevel(aws.LogDebugWithHTTPBody))
		}
		sess = NewAWSSessionWithConfig(*conf)
	} else {
		//Otherwise check for a shared config (e.g. AWS Lambda)
		sess = session.Must(session.NewSessionWithOptions(session.Options{
			SharedConfigState: session.SharedConfigEnable,
		}))
	}
	return sess
}

// NewAWSSessionWithConfig returns a sessions with the given config
func NewAWSSessionWithConfig(config aws.Config) *session.Session {
	return session.Must(session.NewSessionWithOptions(session.Options{
		Config: config,
	}))
}
