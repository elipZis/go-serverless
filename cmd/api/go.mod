module github.com/elipzis/go-serverless/api

go 1.17

require (
	github.com/aws/aws-lambda-go v1.28.0
	github.com/awslabs/aws-lambda-go-api-proxy v0.12.0
	github.com/elipzis/go-serverless v0.0.0
	github.com/go-playground/validator/v10 v10.4.1
	github.com/gofiber/fiber/v2 v2.17.0
	github.com/gofiber/jwt/v2 v2.2.7
	github.com/joho/godotenv v1.4.0
	github.com/kataras/iris/v12 v12.1.8
	github.com/satori/go.uuid v1.2.0
)

require (
	github.com/Shopify/goreferrer v0.0.0-20181106222321-ec9c9a553398 // indirect
	github.com/andybalholm/brotli v1.0.2 // indirect
	github.com/aws/aws-sdk-go v1.43.19 // indirect
	github.com/aymerick/douceur v0.2.0 // indirect
	github.com/fatih/structs v1.1.0 // indirect
	github.com/go-playground/locales v0.13.0 // indirect
	github.com/go-playground/universal-translator v0.17.0 // indirect
	github.com/golang-jwt/jwt/v4 v4.0.0 // indirect
	github.com/gorilla/css v1.0.0 // indirect
	github.com/iris-contrib/blackfriday v2.0.0+incompatible // indirect
	github.com/iris-contrib/schema v0.0.1 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/kataras/golog v0.0.18 // indirect
	github.com/kataras/pio v0.0.8 // indirect
	github.com/klauspost/compress v1.12.2 // indirect
	github.com/leodido/go-urn v1.2.0 // indirect
	github.com/microcosm-cc/bluemonday v1.0.16 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/schollz/closestmatch v2.1.0+incompatible // indirect
	github.com/shurcooL/sanitized_anchor_name v1.0.0 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.26.0 // indirect
	github.com/valyala/tcplisten v1.0.0 // indirect
	golang.org/x/crypto v0.0.0-20210513164829-c07d793c2f9a // indirect
	golang.org/x/net v0.0.0-20220127200216-cd36cc0744dd // indirect
	golang.org/x/sys v0.0.0-20211216021012-1d35b9e2eb4e // indirect
	golang.org/x/text v0.3.7 // indirect
	gopkg.in/yaml.v3 v3.0.0-20200615113413-eeeca48fe776 // indirect
)

replace (
	github.com/elipzis/go-serverless v0.0.0 => ../../
	github.com/elipzis/go-serverless/util v0.0.0 => ../../util
)
