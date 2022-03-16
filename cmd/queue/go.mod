module github.com/elipzis/go-serverless/queue

go 1.17

replace (
	github.com/elipzis/go-serverless v0.0.0 => ../../
	github.com/elipzis/go-serverless/util v0.0.0 => ../../util
)

require (
	github.com/aws/aws-lambda-go v1.28.0
	github.com/aws/aws-sdk-go v1.43.19
	github.com/joho/godotenv v1.4.0
	github.com/satori/go.uuid v1.2.0
)

require github.com/jmespath/go-jmespath v0.4.0 // indirect
