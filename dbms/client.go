package dbms

import(
  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/service/dynamodb"
  "log"
)


type ClientManager struct {
}

type DynamodbConfig struct {
  Region string
  Profile string
}

type DynamodbClient struct{
  Connection *dynamodb.DynamoDB
}

func (cm *ClientManager) Client() *DynamodbClient {
  config := DynamodbConfig{Region: "us-east-1", Profile: ""}
  sess := session.Must(session.NewSessionWithOptions(session.Options{
        Config: aws.Config{Region: aws.String(config.Region)},
        SharedConfigState: session.SharedConfigEnable,
        Profile: config.Profile,
    }))
  return &DynamodbClient{Connection: dynamodb.New(sess)}
}

func (dc *DynamodbClient) GetWorker(key map[string]*dynamodb.AttributeValue)(*dynamodb.GetItemOutput, error){
  input := &dynamodb.GetItemInput{
    Key:       key,
    TableName: aws.String("Workers"),
  }

  result, err := dc.Connection.GetItem(input)
  
  if err != nil {
    return nil, err
  }
  
  return result, nil
}

func (dc *DynamodbClient) CreateWorker(item map[string]*dynamodb.AttributeValue) error {
  input := &dynamodb.PutItemInput{
    Item: item,
    TableName: aws.String("Workers"),
  }

  resp, err := dc.Connection.PutItem(input)

  log.Println(resp)
  return err
}




