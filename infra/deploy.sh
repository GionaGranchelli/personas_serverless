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
  cdk deploy "$STACK_NAME" -c stageName="$STAGE_NAME"
}

# Check if the script is called with an argument
if [ $# -eq 0 ]; then
  echo "Usage: $0 {dev|st|prod}"
  exit 1
fi

# Deploy based on the provided environment
deploy "$1"
