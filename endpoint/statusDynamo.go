package endpoint

import (
	"context"
	"os"
	"github.com/graniticio/granitic/ws"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/graniticio/granitic/logging"
)

type StatusDynamoLogic struct {
	Log logging.Logger
}

func (st *StatusDynamoLogic) Process(ctx context.Context, req *ws.WsRequest, res *ws.WsResponse) {

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("AWS_REGION"))},
	)

	dynamoClient := dynamodb.New(sess)

	input := &dynamodb.DescribeTableInput{
    TableName: aws.String("Workers"),
  }

	result, err := dynamoClient.DescribeTable(input)

	if err != nil {
		st.Log.LogErrorf("Error: %s", err)

		res.Body = "Dynamo got error"
	} else {
		table := result.Table

		st.Log.LogInfof("Info: ", table)

		res.Body = "Dynamo working good"
	}
}
