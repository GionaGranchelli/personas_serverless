#!/usr/bin/env node
import 'source-map-support/register';
import * as cdk from 'aws-cdk-lib';
import {PersonaStackService} from '../lib/persona-stack-service';

const app = new cdk.App();
const env = {
    account: process.env.CDK_DEFAULT_ACCOUNT,
    region: process.env.CDK_DEFAULT_REGION
}
const stageName = app.node.tryGetContext('stageName');

new PersonaStackService(app, 'PersonaStackServiceDev', {
    env: env,
    tableName: 'PersonaTableDev',
    stageName: stageName,
});

new PersonaStackService(app, 'PersonaStackServiceStaging', {
    env: env,
    tableName: 'PersonaTableST',
    stageName: stageName,
});

new PersonaStackService(app, 'PersonaStackService', {
    env: env,
    tableName: 'PersonaTable',
    stageName: stageName,

});