module github.com/elipzis/go-serverless/api

go 1.17

require (
	github.com/aws/aws-lambda-go v1.28.0
	github.com/awslabs/aws-lambda-go-api-proxy v0.12.0
	github.com/elipzis/go-serverless v0.0.0
	github.com/gofiber/fiber/v2 v2.17.0
	github.com/gofiber/jwt/v2 v2.2.7
	github.com/joho/godotenv v1.4.0
)

require (
	github.com/andybalholm/brotli v1.0.2 // indirect
	github.com/aws/aws-sdk-go v1.43.19 // indirect
	github.com/golang-jwt/jwt/v4 v4.0.0 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/klauspost/compress v1.12.2 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.26.0 // indirect
	github.com/valyala/tcplisten v1.0.0 // indirect
	golang.org/x/sys v0.0.0-20211216021012-1d35b9e2eb4e // indirect
)

replace (
	github.com/elipzis/go-serverless v0.0.0 => ../../
	github.com/elipzis/go-serverless/util v0.0.0 => ../../util
)
