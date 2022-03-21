# 🚀 Go Serverless!

An accompanying example project to the [Conf42: Golang 2022 presentation](https://www.conf42.com/Golang_2022_Savas_Ziplies_serverless) by Savas Ziplies.

A scaffold to deploy Golang framework-driven Serverless functions to AWS.

## Pre-requisites

* [AWS CLI](https://aws.amazon.com/cli/) already configured with required permissions
* [Docker installed](https://docs.docker.com/get-docker/)
* [Golang](https://golang.org)
* [SAM CLI](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/serverless-sam-cli-install.html)

## Quickstart

To run/test locally:

* Run `docker-compose up -d timescaledb-postgis localstack`
* (Optional) Create a local queue in localstack for testing: `aws --endpoint-url=http://localhost:4566 sqs create-queue --queue-name goserverless-queue-dev`
* Copy `.env.example` to `.env`
* Run the single commands or deploy

### Deploy

* `cd .build/aws`
* `sam build`
* `sam deploy --guided` for a first time deployment

### Running

* The "api" cmd: `go build github.com/elipzis/go-serverless/api`
* The "queue" cmd: `go build github.com/elipzis/go-serverless/queue`
* The "web" cmd: `go build github.com/elipzis/go-serverless/web`

## Tech Stack

The main tech stack contains of

* [AWS SAM](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/serverless-sam-cli-install.html)
* [Golang](https://golang.org)
* [Docker](https://docs.docker.com/get-docker/)
* [Postgres](https://www.postgresql.org/download/)
* [AWS SQS](https://aws.amazon.com/sqs/)
* [Fiber](https://github.com/gofiber/fiber)
* [AWS Lambda](https://aws.amazon.com/lambda/)
* [Chromium](https://github.com/alixaxel/chrome-aws-lambda)

## Credits

- [elipZis GmbH](https://elipZis.com)
- [NeA](https://github.com/nea)
- [All Contributors](https://github.com/elipZis/go-serverless/contributors)

## Thanks

Thanks to Conf42 for having me :)
- Web - [Conf42](https://www.conf42.com/)
-	LinkedIn - [@conf42](https://www.linkedin.com/company/conf42)
-	Twitter - [@conf42com](https://twitter.com/conf42com)

## License

The MIT License (MIT). Please see [License File](LICENSE.md) for more information.
