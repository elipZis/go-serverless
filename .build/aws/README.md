# Go Serverless! Example AWS Deployment

AWS SAM is the deployment foundation for the example app

## Requirements

* AWS CLI already configured with Administrator permission
* [Docker installed](https://www.docker.com/community-edition)
* [Golang](https://golang.org)
* SAM CLI
    - [Install the SAM CLI](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/serverless-sam-cli-install.html)

## Quickstart

* `sam build`
* `sam deploy --guided` for a first time deployment

### AWS Profiles

If you have a specific profile for the Go Serverless Example App AWS access, add it via e.g. `sam deploy --profile goserverless`.

### AWS SAM Environments

To configure or use a pre-configured `dev` or `prod` AWS SAM config in `samconfig.toml` use
e.g. `sam deploy --profile goserverless --config-env=prod`

### Environment variables

You can deploy to the `dev` and the `prod` environment by passing `--parameter-overrides Environment=dev` or `prod` to
the deployment step.

## Setup process

### Installing dependencies & building the target

Read more
about [SAM Build here](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/sam-cli-command-reference-sam-build.html)

The `sam build` command is to package the serverless app(s). To execute this simply run

```shell
sam build
```

## Packaging and deployment

To deploy your application for the first time, run the following in your shell:

```bash
sam deploy --guided
```

The command will package and deploy your application to AWS, with a series of prompts:

* **Stack Name**: The name of the stack to deploy to CloudFormation. This should be unique to your account and region,
  and a good starting point would be something matching your project name.
* **AWS Region**: The AWS region you want to deploy your app to.
* **Confirm changes before deploy**: If set to yes, any change sets will be shown to you before execution for manual
  review. If set to no, the AWS SAM CLI will automatically deploy application changes.
* **Allow SAM CLI IAM role creation**: Many AWS SAM templates, including this example, create AWS IAM roles required for
  the AWS Lambda function(s) included to access AWS services. By default, these are scoped down to minimum required
  permissions. To deploy an AWS CloudFormation stack which creates or modifies IAM roles, the `CAPABILITY_IAM` value
  for `capabilities` must be provided. If permission isn't provided through this prompt, to deploy this example you must
  explicitly pass `--capabilities CAPABILITY_IAM` to the `sam deploy` command.
* **Save arguments to samconfig.toml**: If set to yes, your choices will be saved to a configuration file inside the
  project, so that in the future you can just re-run `sam deploy` without parameters to deploy changes to your
  application.

You can find your API Gateway Endpoint URL in the output values displayed after deployment.
