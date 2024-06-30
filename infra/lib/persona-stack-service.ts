import * as cdk from 'aws-cdk-lib'
import { Construct } from 'constructs'
import * as dynamodb from 'aws-cdk-lib/aws-dynamodb'
import * as lambda from 'aws-cdk-lib/aws-lambda'
import * as apigateway from 'aws-cdk-lib/aws-apigateway'
import * as events from 'aws-cdk-lib/aws-events'
import * as targets from 'aws-cdk-lib/aws-events-targets'
import * as logs from 'aws-cdk-lib/aws-logs'
import * as iam from 'aws-cdk-lib/aws-iam'
import * as s3 from 'aws-cdk-lib/aws-s3'
import * as s3deploy from 'aws-cdk-lib/aws-s3-deployment'
import * as cloudfront from 'aws-cdk-lib/aws-cloudfront'
import * as cloudfront_origins from 'aws-cdk-lib/aws-cloudfront-origins'


interface PersonaStackServiceProps extends cdk.StackProps {
  tableName: string;
  stageName: string;
}

const apiGatewayOptionsCors = {
  authorizationType: apigateway.AuthorizationType.NONE,
  methodResponses: [{
    statusCode: '200',
    responseParameters: {
      'method.response.header.Access-Control-Allow-Origin': true
    }
  }]
}

export class PersonaStackService extends cdk.Stack {
  constructor(scope: Construct, id: string, props: PersonaStackServiceProps) {
    super(scope, id, props)

    // Creation of a Table in DynamoDB
    const table = new dynamodb.Table(this, 'PersonaTable', {
      partitionKey: { name: 'id', type: dynamodb.AttributeType.STRING },
      tableName: props.tableName,
      billingMode: dynamodb.BillingMode.PAY_PER_REQUEST
    })

    // Setups the Lambda function
    const personaLambda = new lambda.Function(this, 'PersonFunction', {
      runtime: lambda.Runtime.PROVIDED_AL2023,
      code: lambda.Code.fromAsset('../lambda'),
      handler: 'bootstrap',
      environment: {
        TABLE_NAME: table.tableName
      }
    })

    // Grant Lambda permissions to read/write from the table
    table.grantReadWriteData(personaLambda)
    // Grant Lambda permissions to publish events to EventBridge
    personaLambda.addToRolePolicy(new iam.PolicyStatement({
      actions: ['events:PutEvents'],
      resources: ['arn:aws:events:*:*:event-bus/*']
    }))

    // API Gateway
    const api = new apigateway.RestApi(this, 'persona-api', {
      restApiName: 'Persona Service',
      deployOptions: {
        stageName: props.stageName
      },
      defaultCorsPreflightOptions: {
        allowOrigins: apigateway.Cors.ALL_ORIGINS,
        allowMethods: apigateway.Cors.ALL_METHODS, // This is also the default
      },
    })

    const personaResource = api.root.addResource('personas') // Create /personas resource
    personaResource.addMethod('POST', new apigateway.LambdaIntegration(personaLambda), apiGatewayOptionsCors)
    personaResource.addMethod('GET', new apigateway.LambdaIntegration(personaLambda), apiGatewayOptionsCors)

    const personaByIdResource = personaResource.addResource('{id}') // create /personas/{id} resource
    personaByIdResource.addMethod('GET', new apigateway.LambdaIntegration(personaLambda), apiGatewayOptionsCors)
    // personaByIdResource.addMethod('PUT', new apigateway.LambdaIntegration(personaLambda))
    // personaByIdResource.addMethod('DELETE', new apigateway.LambdaIntegration(personaLambda))

    // EventBridge Rule for Person Created event
    const rule = new events.Rule(this, 'PersonCreatedRule', {
      eventPattern: {
        source: ['my.person.service'],
        detailType: ['PersonCreated']
      }
    })
    // Define CloudWatch Log Group for logging EventBridge events
    const logGroup = new logs.LogGroup(this, 'EventBridgeLogGroup', {
      removalPolicy: cdk.RemovalPolicy.DESTROY
    })

    rule.addTarget(new targets.LambdaFunction(personaLambda))
    rule.addTarget(new targets.CloudWatchLogGroup(logGroup))

    // Define S3 bucket for hosting the Vue.js app
    const websiteBucket = new s3.Bucket(this, 'WebsiteBucket', {
      websiteIndexDocument: 'index.html',
      publicReadAccess: true,
      removalPolicy: cdk.RemovalPolicy.DESTROY,
      autoDeleteObjects: true,
      blockPublicAccess: new s3.BlockPublicAccess({
        blockPublicAcls: false,
        ignorePublicAcls: false,
        blockPublicPolicy: false,
        restrictPublicBuckets: false,
      }),
    })
    const distribution = new cloudfront.Distribution(this, 'WebsiteDistribution', {
      defaultBehavior: { origin: new cloudfront_origins.S3Origin(websiteBucket) }
    })
    // Deploy Vue.js build to S3 bucket
    new s3deploy.BucketDeployment(this, 'DeployWebsite', {
      sources: [s3deploy.Source.asset('../persona-app-ts/dist')],
      destinationBucket: websiteBucket,
      distribution,
      distributionPaths: ['/*']
    })
    // Output the CloudFront URL
    new cdk.CfnOutput(this, 'WebsiteURL', {
      value: distribution.distributionDomainName,
    });

    // Output the API Gateway URL
    new cdk.CfnOutput(this, 'ApiGatewayUrl', {
      value: api.url,
    });
  }
}
