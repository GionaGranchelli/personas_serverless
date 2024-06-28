# personas_serverless

# Setup Dependencies

## Install AWS CLI
```@shell
sudo dnf install awscli
```
## Install AWS SAM CLI
```@shell
wget https://github.com/aws/aws-sam-cli/releases/latest/download/aws-sam-cli-linux-x86_64.zip
unzip aws-sam-cli-linux-x86_64.zip -d sam-installation
sudo ./sam-installation/install
```

## Install cdk dependencies
```@shell
npm install @aws-cdk/aws-lambda @aws-cdk/aws-apigateway @aws-cdk/aws-dynamodb @aws-cdk/aws-events-targets
```