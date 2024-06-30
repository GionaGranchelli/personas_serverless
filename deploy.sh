#!/bin/bash

# Function to load environment variables from a file
load_env() {
  if [ -f "$1" ]; then
    export $(cat "$1" | xargs)
  else
    echo "Environment file $1 not found."
    exit 1
  fi
}

# Function to deploy using CDK
deploy() {

  echo "Building Lambdas..."
  (cd lambda && ./build.sh && cd ..) || exit

  echo "Current Location: $(pwd)"

  echo "Building Vue.js application...  "
  (cd persona-app-ts && npm run build && cd ..) || exit

  cd infra || exit
  case $1 in
      dev)
        load_env .env.dev
        STACK_NAME="PersonaStackServiceDev"
        STAGE_NAME="dev"
        ;;
      st)
        load_env .env.st
        STACK_NAME="PersonaStackServiceStaging"
        STAGE_NAME="st"
        ;;
      prod)
        load_env .env.prod
        STACK_NAME="PersonaStackService"
        STAGE_NAME="prod"
        ;;
      *)
        echo "Invalid environment: $1"
        echo "Usage: $0 {dev|st|prod}"
        exit 1
        ;;
  esac

  echo "Deploying to $1 environment..."
  npm run build

  # Get the API Gateway URL from the CloudFormation stack outputs
  API_URL=$(aws cloudformation describe-stacks --stack-name "$STACK_NAME" --query "Stacks[0].Outputs[?OutputKey=='ApiGatewayUrl'].OutputValue" --output text)

  echo "API URL: $API_URL"
  # Update api-config.json with the API URL
  echo "{\"apiUrl\": \"$API_URL\"}" > ../persona-app-ts/public/api-config.json
  echo "Current Location: $(pwd)"
  echo "STACK_NAME: $STACK_NAME"
  echo "STAGE_NAME: $STAGE_NAME"

  cdk deploy "$STACK_NAME" -c stageName="$STAGE_NAME"
}

# Check if the script is called with an argument
if [ $# -eq 0 ]; then
  echo "Usage: $0 {dev|st|prod}"
  exit 1
fi

# Deploy based on the provided environment
deploy "$1"
