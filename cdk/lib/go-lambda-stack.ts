import * as cdk from "@aws-cdk/core";
import * as golang from "aws-lambda-golang";
import * as events from "@aws-cdk/aws-events";
import * as targets from "@aws-cdk/aws-events-targets";
import * as sns from "@aws-cdk/aws-sns";
import * as subscriptions from "@aws-cdk/aws-sns-subscriptions";
import * as iam from "@aws-cdk/aws-iam";
import path = require("path");

export class GoLambdaStack extends cdk.Stack {
  constructor(scope: cdk.Construct, id: string, props?: cdk.StackProps) {
    super(scope, id, props);

    // create a topic where the lambda will publish to
    const topic = new sns.Topic(this, "topic");
    topic.addSubscription(
      new subscriptions.EmailSubscription("email@google.com")
    );

    // create the lambda and let it publish to the above topic
    const lambda = new golang.GolangFunction(this, "ces-lambda");
    lambda.addEnvironment("topic_arn", topic.topicArn);
    lambda.addPermission("sns publish", {
      principal: new iam.ServicePrincipal("sns.amazonaws.com"),
      action: "SNS:Publish",
    });

    // run the lambda once a week
    new events.Rule(this, "weekly-check", {
      schedule: events.Schedule.rate(cdk.Duration.days(7)),
      targets: [new targets.LambdaFunction(lambda)],
    });
  }
}
