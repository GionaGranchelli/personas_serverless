# Persona Management System

The Persona Management System is a serverless application designed to manage personal information efficiently. The
system supports CRUD operations and event processing, leveraging AWS services for scalability, reliability, and reduced
operational overhead.

## Project Composition

### Components

1. **Lambda Functions**
    - **Persona Function**: Handles CRUD operations for persona data.
    - **Event API Function**: Provides an API to fetch stored events from DynamoDB.

2. **Amazon API Gateway**
    - Provides secure and scalable entry points for API requests.
    - Routes requests to the appropriate Lambda functions.

3. **Amazon DynamoDB**
    - **PersonaTable**: Stores persona data.
    - **EventTable**: Stores event data triggered by CRUD operations.

4. **Amazon EventBridge**
    - Manages event-driven workflows for persona creation, update, and deletion.

5. **Amazon S3 and CloudFront**
    - Hosts the Vue.js frontend application and delivers it efficiently to users.

6. **Vue.js Frontend**
    - Provides a user-friendly interface for managing personas.
    - Displays a list of personas and event logs.

## Design Choices

- **Serverless Architecture**: Leveraging AWS Lambda, API Gateway, and DynamoDB to ensure scalability and reliability
  without managing infrastructure.
- **Event-Driven**: Using EventBridge to handle event-driven workflows, ensuring that updates and deletions trigger
  appropriate actions.
- **Modular Components**: Separating the application into distinct Lambda functions for better maintainability and
  scalability.
- **Frontend Framework**: Using Vue.js for a responsive and interactive user interface.

## Running the Project

### Prerequisites

- **Node.js** and **npm**
- **AWS CLI**
- **AWS CDK**
- **Docker** (for SAM local testing)

### Setup

## Install AWS CLI

For Linux:

```@shell
sudo dnf install awscli
```

For the rest of the OS:

```@shell
npm i aws-cli
```

## Install AWS SAM CLI

```@shell
wget https://github.com/aws/aws-sam-cli/releases/latest/download/aws-sam-cli-linux-x86_64.zip
unzip aws-sam-cli-linux-x86_64.zip -d sam-installation
sudo ./sam-installation/install
```

## Install cdk dependencies
Install dependencies for CDK:
```@shell
   cd infra
   npm install
```

Install dependencies for Frontend:
```sh
   cd ../persona-app-ts
   npm install
```

## Running with AWS CDK
If you want to deploy the application using AWS CDK, you can do so by running the following commands:
```sh
cd infra
cdk bootstrap
cdk deploy
```

## Running with SAM Local
If you want to run the Lambda functions locally using SAM, you can do so by running the following commands:
```sh
sam build
sam local start-api
```

## Deployment Script To deploy to different environments

The deploy.sh script automates the deployment process for the Persona Management System. It performs the following
tasks:

##### Loads Environment Variables: 
Depending on the selected environment:
- **development** 
- **staging**
- **production**

The script loads the appropriate environment variables from .env.development, .env.staging, or .env.production files ( Located in the infra directory ).
- Build the Lambda functions: It navigates to the Lambda functions directory and runs the build process to generate the
  deployment packages for the Go Project.
- Builds the Vue.js Application: It navigates to the Vue.js project directory and runs the build process to generate the
  production-ready frontend assets.
- Deploys AWS Infrastructure: Using AWS CDK, the script builds and deploys the infrastructure defined in the infra
  directory.
- Updates API Configuration: After deployment, it retrieves the API Gateway URL from the AWS CloudFormation stack
  outputs and updates the api-config.json file in the Vue.js project with the new API endpoint.
- Deploys the CDK Stack: Finally, the script deploys the CDK stack to the specified environment.

To deploy the application, run the script with one of the following arguments:

```sh
./deploy.sh dev    # Deploy to the development environment
./deploy.sh st     # Deploy to the staging environment
./deploy.sh prod   # Deploy to the production environment
```

This script ensures a consistent and automated deployment process across different environments.

# Well known issue
- The deployment script is not tested on Windows OS. It is tested only Linux OS - Fedora 40.





