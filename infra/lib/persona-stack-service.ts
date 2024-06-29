import * as cdk from 'aws-cdk-lib';
import { Construct } from 'constructs';
import * as dynamodb from 'aws-cdk-lib/aws-dynamodb';
import * as lambda from 'aws-cdk-lib/aws-lambda';
import * as apigateway from 'aws-cdk-lib/aws-apigateway';
import * as events from 'aws-cdk-lib/aws-events';
import * as targets from 'aws-cdk-lib/aws-events-targets';
import * as logs from 'aws-cdk-lib/aws-logs';
import * as iam from 'aws-cdk-lib/aws-iam';


export class PersonaStackService extends cdk.Stack {
  constructor(scope: Construct, id: string, props?: cdk.StackProps) {
    super(scope, id, props);

// Creation of a Table in DynamoDB
    const table = new dynamodb.Table(this, 'PersonaTable', {
      partitionKey: { name: 'id', type: dynamodb.AttributeType.STRING },
      tableName: 'PersonaTable',
      billingMode: dynamodb.BillingMode.PAY_PER_REQUEST,
    });

    // Setups the Lambda function
    const personaLambda = new lambda.Function(this, 'PersonFunction', {
      runtime: lambda.Runtime.PROVIDED_AL2023,
      code: lambda.Code.fromAsset('../lambda'),
      handler: 'bootstrap',
      environment: {
        TABLE_NAME: table.tableName
      }
    });

    // Grant Lambda permissions to read/write from the table
    table.grantReadWriteData(personaLambda);
    // Grant Lambda permissions to publish events to EventBridge
    personaLambda.addToRolePolicy(new iam.PolicyStatement({
      actions: ['events:PutEvents'],
      resources: ['arn:aws:events:*:*:event-bus/*'],
    }));

    // API Gateway
    const api = new apigateway.RestApi(this, 'persona-api', {
      restApiName: 'Persona Service'
    });

    const personaResource = api.root.addResource('personas');
    personaResource.addMethod('POST', new apigateway.LambdaIntegration(personaLambda));

    // EventBridge Rule for Person Created event
    const rule = new events.Rule(this, 'PersonCreatedRule', {
      eventPattern: {
        source: ['my.person.service'],
        detailType: ['PersonCreated']
      }
    });
    // Define CloudWatch Log Group for logging EventBridge events
    const logGroup = new logs.LogGroup(this, 'EventBridgeLogGroup', {
      removalPolicy: cdk.RemovalPolicy.DESTROY,
    });

    rule.addTarget(new targets.LambdaFunction(personaLambda));
    rule.addTarget(new targets.CloudWatchLogGroup(logGroup));
  }
}
