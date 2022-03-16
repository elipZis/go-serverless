module github.com/elipzis/go-serverless/web

go 1.17

require (
	github.com/aws/aws-lambda-go v1.28.0
	github.com/awslabs/aws-lambda-go-api-proxy v0.12.0
	github.com/elipzis/go-serverless v0.0.0
	github.com/gobuffalo/packr/v2 v2.8.3
	github.com/gofiber/fiber/v2 v2.29.0
	github.com/gofiber/template v1.6.25
	github.com/joho/godotenv v1.4.0
)

require (
	github.com/GeertJohan/go.incremental v1.0.0 // indirect
	github.com/GeertJohan/go.rice v1.0.2 // indirect
	github.com/akavel/rsrc v0.8.0 // indirect
	github.com/andybalholm/brotli v1.0.4 // indirect
	github.com/aws/aws-sdk-go v1.43.19 // indirect
	github.com/daaku/go.zipexe v1.0.0 // indirect
	github.com/gobuffalo/logger v1.0.6 // indirect
	github.com/gobuffalo/packd v1.0.1 // indirect
	github.com/jessevdk/go-flags v1.4.0 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/karrick/godirwalk v1.16.1 // indirect
	github.com/klauspost/compress v1.15.0 // indirect
	github.com/markbates/errx v1.1.0 // indirect
	github.com/markbates/oncer v1.0.0 // indirect
	github.com/markbates/safe v1.0.1 // indirect
	github.com/nkovacs/streamquote v1.0.0 // indirect
	github.com/sirupsen/logrus v1.8.1 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.34.0 // indirect
	github.com/valyala/fasttemplate v1.2.1 // indirect
	github.com/valyala/tcplisten v1.0.0 // indirect
	golang.org/x/sys v0.0.0-20220227234510-4e6760a101f9 // indirect
	golang.org/x/term v0.0.0-20210927222741-03fcf44c2211 // indirect
)

replace (
	github.com/elipzis/go-serverless v0.0.0 => ../../
	github.com/elipzis/go-serverless/util v0.0.0 => ../../util
)
