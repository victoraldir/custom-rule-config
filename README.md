# AWS Config custom rule

This repository contains a custom rule for AWS Config that checks if the following set of tags are present on the resources:

- `Environment`
- `Owner`
- `Application`

The rule is written in Go and can be found withing the [required-tags](/required-tags/) directory.

## Remediation

The rule is non-compliant if any of the tags are missing. The remediation is to add the missing tags to the resource. This is also done via a Lambda function written in Go and can be found within the [required-tags-remediation](/required-tags-remediation/) directory.

## Deployment

This project uses SAM to deploy the Lambda functions. To deploy it, make sure you have the AWS CLI and SAM CLI installed and configured. Then, run the following commands:

```bash
sam build
sam deploy --guided
```

This will guide you through the deployment process.

> **_NOTE:_** 